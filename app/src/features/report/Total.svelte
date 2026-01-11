<script lang="ts">
  import { getTotal } from "./report.svelte";
  import { onMount } from "svelte";

  let total = $state(0);

  onMount(async () => {
    const res = await getTotal();
    console.log(res);
    total = res;
  });
</script>

<article>
  <header>総収支</header>
  <span>{total} 円</span>
</article>
