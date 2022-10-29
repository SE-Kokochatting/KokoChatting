/**
 * description: 主题颜色状态管理
 * author: Yuming Cui
 * date: 2022-10-28 23:21:28 +0800
 */

import { makeAutoObservable } from 'mobx'
import { Theme } from '@/enums'

class ThemeState {
  public theme = Theme.Light

  public constructor() {
    this.init()
    makeAutoObservable(this)
  }

  /**
   * 获取主题
   * @param void
   * @returns 主题
   */
  public getTheme(): Theme | null {
    return localStorage.getItem('theme') as Theme | null
  }

  /**
   * 设置主题
   * @param val 要设置的值
   * @returns void
   */
  public setTheme(val: Theme) {
    this.theme = val
    localStorage.setItem('theme', val)
  }

  /**
   * 初始化
   * @param void
   * @returns void
   */
  private init() {
    const themeValue = this.getTheme()
    if (themeValue === null || themeValue === Theme.Light) {
      this.setTheme(Theme.Light)
      this.theme = Theme.Light
    } else {
      this.theme = Theme.Dark
    }
  }
}

const ThemeStore = new ThemeState()

export default ThemeStore
