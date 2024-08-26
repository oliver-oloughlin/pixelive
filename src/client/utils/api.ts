import { jex } from "@olli/jex";
import { PixelSchema } from "./models";
import { logger } from "@olli/jex/logger";

export const api = jex({
  baseUrl: "/api",
  plugins: [logger()],
  endpoints: {
    pixels: {
      get: {
        data: PixelSchema.array(),
      },
      post: {
        body: PixelSchema,
      },
    },
  },
});
