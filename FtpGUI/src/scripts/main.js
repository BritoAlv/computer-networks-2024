import { connectionSetupComponent } from "./components/connectionSetup.js";
import { connectionStatusComponent } from "./components/connectionStatus.js";
import { localDirectoryComponent } from "./components/localDirectory.js";
import { serverDirectoryComponent } from "./components/serverDirectory.js";

// Load required components
connectionSetupComponent();
connectionStatusComponent();
await localDirectoryComponent();
serverDirectoryComponent();