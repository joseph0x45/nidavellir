<script>
  import { onMount } from "svelte";
  import Package from "../components/Package.svelte";
  import Breadcrumb from "../components/Breadcrumb.svelte";

  /**
   * @typedef {Object} Package
   * @property {string} id
   * @property {string} name
   * @property {string} description
   * @property {string} package_type
   * @property {string} repo_url
   */
  /** @type {Package[]} */
  let packages = [];
  let loading = true;
  let error = "";

  onMount(async () => {
    try {
      const response = await fetch("/api/packages");
      if (!response.ok) throw new Error(`HTTP ${response.status}`);
      const data = await response.json();
      packages = data.data;
    } catch (err) {
      error = "Failed to fetch packages. Please try again.";
      console.error(err);
    } finally {
      loading = false;
    }
  });
</script>

<Breadcrumb items={[{ label: "Packages", href: "/" }]} />
<main class="p-8 max-w-4xl mx-auto">
  {#if loading}
    <p>Loading packages...</p>
  {:else if error}
    <p class="text-red-600">{error}</p>
  {:else if packages.length === 0}
    <p>No packages found.</p>
  {:else}
    <ul class="space-y-4">
      {#each packages as pkg}
        <li>
          <Package {pkg} />
        </li>
      {/each}
    </ul>
  {/if}
</main>
