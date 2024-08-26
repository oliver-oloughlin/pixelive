import { Pixel, PixelSchema } from "./models";
import { setPixel } from "./store";

class ApiWebSocket {
  private attempts: number;
  private ws: WebSocket | null;
  public onSetPixel: ((pixel: Pixel) => unknown) | null;

  constructor(attempts?: number) {
    this.attempts = Math.max(1, attempts ?? 1);
    this.onSetPixel = null;
    this.ws = this.createWs();
  }

  setPixel(pixel: Pixel) {
    this.ws?.send(JSON.stringify(pixel));
  }

  close() {
    this.ws?.close();
  }

  private createWs() {
    if (this.attempts <= 0) {
      return null;
    }

    this.attempts--;
    const ws = new WebSocket("ws://localhost:8080/api/ws");

    ws.onmessage = async (msg) => {
      try {
        const data = JSON.parse(msg.data);
        const parsed = PixelSchema.parse(data);
        await this.onSetPixel?.(parsed);
      } catch (e) {
        console.error(e);
      }
    };

    ws.onerror = (e) => {
      console.error(e);
      this.ws = this.createWs();
    };

    ws.onclose = (e) => {
      console.log(e);
      this.ws = this.createWs();
    };

    return ws;
  }
}

const apiWs = new ApiWebSocket();
apiWs.onSetPixel = setPixel;

export { apiWs };
