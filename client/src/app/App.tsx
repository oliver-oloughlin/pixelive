import { ColorPicker } from "./ColorPicker";
import { DrawingCanvas } from "./DrawingCanvas";

export function App() {
  return (
    <main class="w-[100dvw] h-[100dvh] place-content-center grid">
      <DrawingCanvas />
      <ColorPicker />
    </main>
  );
}
