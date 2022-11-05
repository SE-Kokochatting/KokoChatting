/**
 * description: 订阅发布
 * author: Yuming Cui
 * date: 2022-11-05 14:20:59 +0800
 */

import EventEmitter from 'eventemitter3'

const eventEmitter = new EventEmitter()

const Emitter = {
  on: (event: string, fn: () => void) => eventEmitter.on(event, fn),
  once: (event: string, fn: () => void) => eventEmitter.once(event, fn),
  off: (event: string, fn: () => void) => eventEmitter.off(event, fn),
  emit: (event: string, payload?: any) => eventEmitter.emit(event, payload),
  removeListener: (event: string) => eventEmitter.removeListener(event),
}

Object.freeze(Emitter)

export default Emitter
