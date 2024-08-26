import { DrawPixel } from "./DrawPixel";
import { store } from "../utils/store";

export function DrawingCanvas() {
  return (
    <div class="w-[95vmin] aspect-square border-2 border-slate-400 grid grid-cols-[repeat(100,1fr)]">
      {sortedPixels().map(([id, signal]) => <DrawPixel id={id} signal={signal} />)}
    </div>
  )
}

function sortedPixels() {
  const [pixels] = store.pixelsResource
  return Array.from(pixels()?.entries() ?? [])
    .sort(([a], [b]) => a - b)
}
