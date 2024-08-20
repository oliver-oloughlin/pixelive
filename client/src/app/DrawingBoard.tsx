import { DrawPixel } from "./Pixel";
import { store } from "../utils/store";

export function DrawingBoard() {
  return (
    <div class="w-[95vmin] aspect-square border-2 border-slate-400 grid grid-cols-[repeat(100,1fr)]">
      {store.pixels.map(signal => <DrawPixel signal={signal} />)}
    </div>
  )
}
