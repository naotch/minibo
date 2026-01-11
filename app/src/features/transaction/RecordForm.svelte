<script lang="ts">
  import { record } from "./transaction.svelte";

  let title = $state("");
  let category = $state(0);
  let amount = $state(0);

  function onclick() {
    record(title, category, amount);
  }
</script>

<article>
  <header>収支の記録</header>
  <label for="title">タイトル</label>
  <input type="text" id="title" bind:value={title} />
  <label for="amount">金額</label>
  <input type="number" id="amount" bind:value={amount} />
  <fieldset>
    <legend>カテゴリー</legend>
    <input type="radio" id="outgo" value={0} bind:group={category} />
    <label for="outgo">支出</label>
    <input type="radio" id="income" value={1} bind:group={category} />
    <label for="income">収入</label>
  </fieldset>
  <button class="primary" {onclick}>記録</button>
</article>
