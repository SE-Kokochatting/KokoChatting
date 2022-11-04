/**
 * description: 用户信息状态管理
 * author: Yuming Cui
 * date: 2022-10-28 23:21:42 +0800
 */

import { makeAutoObservable } from 'mobx'

// 用于渲染右侧“用户信息”
class UserState {
  // 是否在右侧展示用户信息
  public showUserInfo = false
  // uid
  public uid = 0
  // 昵称
  public name = ''
  // 头像存储路径
  public avatarUrl = ''

  public constructor() {
    makeAutoObservable(this)
  }

  /**
   * 设置显示状态
   * @param val 要设置的值
   * @returns void
   */
  public setShowUserInfo(val: boolean) {
    this.showUserInfo = val
  }

  /**
   * 保存用户信息
   * @param
   * @returns void
   */
  public setUserInfo({ uid, name, avatarUrl }: any) {
    this.uid = uid
    this.name = name
    this.avatarUrl = avatarUrl
  }

}

const UserStore = new UserState()

export default UserStore
