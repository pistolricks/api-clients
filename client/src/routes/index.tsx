import { A } from "@solidjs/router";
import Counter from "~/components/Counter";
import MapComponent from "~/components/SchoolApiMap";
import BaseDrawer from "~/components/ui/drawer/drawer";

export default function Home() {
  return (
    <main class="text-center mx-auto text-gray-700">

          <MapComponent />

    </main>
  );
}
