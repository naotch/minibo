<script lang="ts">
  import { modal } from "../../stores/store.svelte";
  let { onclick } = $props();
</script>

<dialog open>
  <article>
    <header><div>{modal.title}</div></header>
    <p>{modal.message}</p>
    <footer>
      <button {onclick}>閉じる</button>
    </footer>
  </article>
</dialog>
