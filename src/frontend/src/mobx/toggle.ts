/**
 * description: toggle 状态管理
 * author: Yuming Cui
 * date: 2022-10-31 16:19:44 +0800
 */

import { makeAutoObservable } from 'mobx'
import { ToggleType } from '@/enums'

class ToggleState {
  // 是否显示 toggle
  public showToggle = false
  // toggle 类型
  public toggleType: ToggleType = ToggleType.CreateGroup

  public constructor() {
    makeAutoObservable(this)
  }

  /**
   * 设置显示状态
   * @param val 要设置的值
   * @returns void
   */
  public setShowToggle(val: boolean) {
    this.showToggle = val
  }

  /**
   * 设置 toggle 类型
   * @param val 要设置的值
   * @returns void
   */
  public setToggleType(val: ToggleType) {
    this.toggleType = val
  }
}

const ToggleStore = new ToggleState()

export default ToggleStore
