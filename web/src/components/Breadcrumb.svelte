<script>
  /**
   * @typedef {{ label: string, href: string }} Crumb
   */
  /** @type {Crumb[]} */
  export let items = [];
</script>

<nav class="mb-6 text-gray-400">
  {#each items as item, i}
    {#if i > 0}
      <span class="mx-2">›</span>
    {/if}
    <a href={item.href} class="hover:text-gray-200">
      {item.label}
    </a>
  {/each}
</nav>
