import { Pixel, PixelSchema } from "./models"
import { setPixel } from "./store"

class ApiWebSocket {
  private retries: number
  private ws: WebSocket | null
  public onSetPixel: ((pixel: Pixel) => unknown) | null

  constructor() {
    this.retries = 10
    this.onSetPixel = null
    this.ws = this.createWs()
  }

  setPixel(pixel: Pixel) {
    this.ws?.send(JSON.stringify(pixel))
  }

  close() {
    this.ws?.close()
  }

  private createWs() {
    if (this.retries <= 0) return null
    this.retries--

    const ws = new WebSocket("ws://localhost:8080/ws")

    ws.onmessage = async (msg) => {
      try {
        const data = JSON.parse(msg.data)
        const parsed = PixelSchema.parse(data)
        await this.onSetPixel?.(parsed)
      } catch (e) {
        console.error(e)
      }
    }

    ws.onerror = (e) => {
      console.error(e)
      this.ws = this.createWs()
    }

    ws.onclose = (e) => {
      console.log(e)
      this.ws = this.createWs()
    }

    return ws
  }
}

const apiWs = new ApiWebSocket()
apiWs.onSetPixel = setPixel

export { apiWs }
