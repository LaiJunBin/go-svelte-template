<script>
  import { Link } from 'svelte-routing'

  export let to
</script>

<Link {to} class="p-3"><slot /></Link>
