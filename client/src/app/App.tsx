import { ColorPicker } from "./ColorPicker";
import { DrawingBoard } from "./DrawingCanvas";

export function App() {
  return (
    <main class="w-[100dvw] h-[100dvh] place-content-center grid">
      <DrawingBoard />
      <ColorPicker />
    </main>
  );
}
