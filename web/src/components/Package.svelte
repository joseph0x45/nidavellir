<script>
  import { Link } from "svelte-routing";

  /**
   * @typedef {Object} Package
   * @property {string} id
   * @property {string} name
   * @property {string} description
   * @property {string} package_type
   * @property {string} repo_url
   */
  /** @type {Package} */
  export let pkg;
</script>

<Link to={`/packages/${pkg.id}`} class="block p-4 border rounded hover:shadow">
  <h2 class="text-xl font-semibold">{pkg.name}</h2>
  <p class="text-gray-600 mb-2">{pkg.description}</p>
  <p class="text-gray-500"><strong>Type:</strong> {pkg.package_type}</p>
  <p class="text-gray-500">
    <strong>Repo:</strong>
    <a class="text-blue-600 hover:underline" href={pkg.repo_url} target="_blank"
      >{pkg.repo_url}</a
    >
  </p>
</Link>
