// @ts-nocheck

/* eslint-disable */

/**
 * description: 订阅发布
 * author: Others
 * date: 2022-11-05 14:20:59 +0800
 */

export class MessageCenter {
  events = {}
  /**
   * 注册事件至调度中心
   * @param type 事件类型，特指具体事件名
   * @param handler 事件注册的回调
   */
  on(
    type: string | number,
    handler: { (...args: any[]): any; (...args: any[]): void },
  ) {
    // 订阅者
    this.checkHandler(type, handler)
    if (!this.has(type)) {
      // 若调度中心未找到该事件的队列，则新建某个事件列表（可以对某个类型的事件注册多个回调函数）
      this.events[type] = []
    }
    this.events[type].push(handler)
    return this
  }
  /**
   * 触发调度中心的某个或者某些该事件类型下注册的函数
   * @param type 事件类型，特指具体事件名
   * @param data 发布者传递的参数
   */
  emit(type: string, data: any) {
    // 发布者
    if (this.has(type)) {
      this.runHandler(type, data)
    }
    return this
  }
  // 销毁监听
  un(type: any, handler: (...args: any[]) => any) {
    this.unHandler(type, handler)
    return this
  }
  // 只注册一次监听，执行即销毁
  once(type: string, handler: { (value: unknown): void; (arg0: any): any }) {
    this.checkHandler(type, handler)
    const fn = (...args: any[]) => {
      this.un(type, fn)
      return handler(...args)
    }
    this.on(type, fn)
    return this
  }
  // 重置调度中心
  clear() {
    this.events = {}
    return this
  }
  // 判断事件是否被订阅
  has(type: string | number) {
    return !!this.events[type]
  }
  // 同一个事件被绑定了多少函数
  handlerLength(type: string | number) {
    return this.events[type]?.length ?? 0
  }
  // 监听invoke的消息，若handler中进行了计算或者异步操作，会反馈给invoke
  watch(type: any, handler: (arg0: any) => any) {
    this.checkHandler(type, handler)
    const fn = (...args: any[]) => {
      this.emit(this.prefixStr(type), handler(...args))
    }
    this.on(type, fn)
    return this
  }
  // 触发watch事件，并且接收watch处理结果
  invoke(type: any, data: any) {
    return new Promise((resolve) => {
      this.once(this.prefixStr(type), resolve)
      this.emit(type, data)
    })
  }
  // 批量执行调度中心中某个函数集
  runHandler(type: string | number, data: any) {
    for (let i = 0; i < this.events[type].length; i++) {
      this.events[type][i] && this.events[type][i](data)
    }
  }
  // 批量销毁调度中心中某个函数集
  unHandler(type: string | number, handler: any) {
    !handler && (this.events[type] = [])
    handler && this.checkHandler(type, handler)
    for (let i = 0; i < this.events[type].length; i++) {
      if (this.events[type][i] && this.events[type][i] === handler) {
        this.events[type][i] = null
      }
    }
  }
  prefixStr(str: any) {
    return `@${str}`
  }
  /**
   * 检查参数是否符合标准
   * @param type 事件名
   * @param handler 事件钩子
   */
  checkHandler(type: string | any[], handler: any) {
    if (type?.length === 0) {
      throw new Error('type.length can not be 0')
    }
    if (!handler || !type) {
      throw new ReferenceError('type or handler is not defined')
    }
    if (typeof handler !== 'function' || typeof type !== 'string') {
      throw new TypeError(
        `${handler} is not a function or ${type} is not a string`,
      )
    }
  }
  // 返回当前类的实例的单例
  static Instance(Fn: typeof MessageCenter) {
    if (!Fn._instance) {
      Object.defineProperty(Fn, '_instance', {
        value: new Fn(),
      })
    }
    return Fn._instance
  }
}

export const messageCenter = MessageCenter.Instance(MessageCenter)
export default MessageCenter
