/**
 * description: websocket
 * author: Yuming Cui
 * date: 2022-11-01 18:30:15 +0800
 */

import { MessageType } from '@/enums'
import { IMessage } from '@/types'
import Emitter from '@/utils/eventEmitter'
import MsgStore from '@/mobx/msg'

enum ModeCode {
  Msg,
  HeartBeat,
}

interface HeartBeatConfig {
  // 心跳时间间隔
  time: number
  // 心跳超时间隔
  timeout: number
  // 断线重连时间
  reconnect: number
}

export default class WS extends WebSocket {
  private heartBeat: any
  private isReconnect: any
  private reconnectTimer: any
  private waitingTimer: any
  private heartTimer: any
  private webSocketState: any

  public constructor(url: string, protocals?: string[]) {
    super(url, protocals)
  }

  public init(heartBeatConfig: HeartBeatConfig, isReconnect: boolean) {
    this.onopen = this.openHandler // 连接上时回调
    this.onclose = this.closeHandler // 断开连接时回调
    this.onmessage = this.messageHandler // 收到服务端消息
    this.onerror = this.errorHandler // 连接出错
    this.heartBeat = heartBeatConfig
    this.isReconnect = isReconnect
    this.reconnectTimer = null // 断线重连时间器
    this.waitingTimer = null // 超时等待时间器
    this.heartTimer = null // 心跳时间器
    this.webSocketState = false // socket状态 true为已连接
  }

  public openHandler() {
    this.webSocketState = true // socket状态设置为连接，做为后面的断线重连的拦截器
    !!this.heartBeat &&
      !!this.heartBeat.time &&
      this.startHeartBeat(this.heartBeat.time) // 是否启动心跳机制
    console.log('websocket开启')
  }

  public messageHandler(e: any) {
    const msg = this.getMsg(e)
    let realMsgId = 0
    // 接受到数据，根据 MsgType，把对象 push 进 MsgStore 的不同消息数组中
    switch (msg.messageType) {
      case MessageType.PongMessage:
        this.webSocketState = true
        console.log('收到心跳响应' + msg.Time)
        break
      case MessageType.SingleMessage:
        MsgStore.sendMsg(msg, MessageType.SingleMessage)
        // 由于有新的消息，需要让 IntersectionObserver 能够监测到新的 DOM 元素
        Emitter.emit('updateIntersect')
        break
      case MessageType.GroupMessage:
        MsgStore.sendMsg(msg, MessageType.GroupMessage)
        Emitter.emit('updateIntersect')
        break
      case MessageType.FriendRequestNotify:
        MsgStore.sendMsg(msg, MessageType.FriendRequestNotify)
        break
      case MessageType.JoinGroupNotify:
        MsgStore.sendMsg(msg, MessageType.JoinGroupNotify)
        break
      case MessageType.HasReadSingleNotify:
        realMsgId = JSON.parse(msg.messageContent).ReadMsgId
        MsgStore.setMsgRead(MessageType.SingleMessage, realMsgId, [
          msg.senderId,
        ])
        break
      case MessageType.HasReadGroupNotify:
        realMsgId = JSON.parse(msg.messageContent).ReadMsgId
        MsgStore.setMsgRead(MessageType.GroupMessage, realMsgId, [msg.senderId])
        break
    }
  }

  public closeHandler(e: any) {
    this.webSocketState = false
    console.log('websocket断开: ' + e.code + ' ' + e.reason + ' ' + e.wasClean)
  }

  public errorHandler() {
    this.webSocketState = false
    this.reconnectWebSocket()
    console.error('websocket出错')
  }

  public sendMsg(sendData: any) {
    this.send(JSON.stringify(sendData))
  }

  public getMsg(e: any): IMessage & { Time: string } {
    return JSON.parse(e.data)
  }

  // 心跳初始函数
  public startHeartBeat(time: number) {
    this.heartTimer = setTimeout(() => {
      this.sendMsg({
        modeCode: ModeCode.HeartBeat,
        time: new Date(),
      })
      console.log('已发送心跳包')
      this.waitingTimer = this.waitingServer()
    }, time)
  }

  // 延时等待服务端响应，通过 webSocketState 判断是否连线成功
  public waitingServer() {
    this.webSocketState = false
    return setTimeout(() => {
      if (this.webSocketState) return this.startHeartBeat(this.heartBeat.time)
      console.log('心跳无响应，已断线')
      this.reconnectTimer = this.reconnectWebSocket()
    }, this.heartBeat.timeout)
  }

  // 重连操作
  public reconnectWebSocket() {
    if (!this.isReconnect) return
    return setTimeout(() => {
      Emitter.emit('reconnect')
      console.log('已重连')
    }, this.heartBeat.reconnect)
  }
  // 清除所有定时器
  public clearTimer() {
    clearTimeout(this.reconnectTimer)
    clearTimeout(this.heartTimer)
    clearTimeout(this.waitingTimer)
  }
  // 关闭连接
  public clear(isReconnect = false) {
    this.isReconnect = isReconnect
    this.clearTimer()
    this.close()
  }
}
