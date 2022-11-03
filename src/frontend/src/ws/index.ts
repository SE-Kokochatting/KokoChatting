/**
 * description: websocket
 * author: Yuming Cui
 * date: 2022-11-01 18:30:15 +0800
 */

import { MessageType } from '@/enums'

export default class _WebSocket extends WebSocket {
  private constructor(url: string, protocals?: string[]) {
    super(url, protocals)
  }
}
