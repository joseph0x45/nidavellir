<script>
  import "./app.css";
  import { Router, Route, Link } from "svelte-routing";
  import Home from "./pages/home.svelte";
  import PackagePage from "./pages/package_page.svelte";
</script>

<Router>
  <Route path="/"><Home /></Route>
  <Route path="/packages"><PackagePage /></Route>
</Router>
