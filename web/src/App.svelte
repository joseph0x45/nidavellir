<script>
	import './app.css';
	import { Router, Route, Link } from "svelte-routing";
	import Home from "./pages/home.svelte";
</script>

<Router>
  <Route path="/"><Home /></Route>
</Router>
