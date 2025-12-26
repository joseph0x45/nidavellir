<script>
  import "./app.css";
  import { Router, Route } from "svelte-routing";
  import Home from "./pages/home.svelte";
  import PackagePage from "./pages/package_page.svelte";
</script>

<div class="min-h-screen bg-black text-gray-100">
  <!-- Global header -->
  <header class="py-6 border-b border-gray-800">
    <h1 class="text-3xl font-bold text-center">Nidavellir</h1>
  </header>

  <!-- Routed content -->
  <main class="max-w-4xl mx-auto px-6 py-6">
    <Router>
      <Route path="/"><Home /></Route>
      <Route let:params path="/packages/:id">
        <PackagePage packageID={params.id} />
      </Route>
    </Router>
  </main>
</div>
