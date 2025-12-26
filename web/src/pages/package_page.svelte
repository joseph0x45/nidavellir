<script>
  import { onMount } from "svelte";
  import Breadcrumb from "../components/Breadcrumb.svelte";

  let pkg = null;
  let loading = true;
  let error = "";

  /**
   * @type {string}
   */
  export let packageID;

  onMount(async () => {
    try {
      const res = await fetch(`/api/packages/${packageID}`);
      if (!res.ok) throw new Error(`HTTP ${res.status}`);
      const data = await res.json();
      pkg = data.data;
    } catch (err) {
      console.error(err);
      error = "Failed to load package";
    } finally {
      loading = false;
    }
  });
</script>

{#if loading}
  <p>Loading package…</p>
{:else if error}
  <p class="text-red-600">{error}</p>
{:else if pkg}
  <Breadcrumb
    items={[
      { label: "Packages", href: "/" },
      { label: pkg.name, href: `/packages/${pkg.id}` },
    ]}
  />

  <h2 class="text-2xl font-semibold mb-2">{pkg.name}</h2>
  <p class="text-gray-400 mb-6">{pkg.description}</p>

  {#if pkg.releases.length === 0}
    <p class="text-gray-400 italic">No releases available for this package.</p>
  {:else}
    <div class="space-y-4">
      {#each pkg.releases as release}
        <div class="border border-gray-800 rounded p-4">
          <h3 class="font-semibold mb-2">Version {release.version}</h3>

          <ul class="ml-4 space-y-1">
            {#each release.artifacts as artifact}
              <li>
                <a
                  href={artifact.download_url}
                  target="_blank"
                  class="text-blue-400 hover:underline"
                >
                  {artifact.artifact_type}
                </a>
              </li>
            {/each}
          </ul>
        </div>
      {/each}
    </div>
  {/if}
{/if}
