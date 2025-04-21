import {defineConfig} from "@solidjs/start/config";
import tailwindcss from "@tailwindcss/vite";

export default defineConfig({
    ssr: false,
    vite: {
        server: {
            proxy: {
                '/api': {
                    target: 'http://localhost:8080',
                    changeOrigin: true,
                },
            },
        },
        plugins: [tailwindcss()]

    }
});
