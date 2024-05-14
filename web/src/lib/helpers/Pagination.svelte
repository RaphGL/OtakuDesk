<!-- todo: improve how pages are generated -->
<script lang="ts">
  // one should bind to currPage to know what page to render
  type Props = { currPage: number, maxPage: number };
  let { currPage = $bindable(1), maxPage  = 10}: Props = $props();

  let pages = [];
  for (let i = currPage; i <= maxPage; i++) {
    pages.push(i);
  }
</script>

<div class="is-flex is-justify-content-center">
  <nav class="pagination" aria-label="pagination">
    <button onclick={() => --currPage} class="pagination-previous">Prev</button>
    <button onclick={() => ++currPage} class="pagination-next">Next</button>

    <ul class="pagination-list">
      {#if currPage > maxPage}
        <li>
          <a href="#" class="pagination-link" aria-label="Goto page 1">1</a>
        </li>
        <li>
          <span class="pagination-ellipsis">&hellip;</span>
        </li>
      {/if}

      {#each pages as page}
        <li>
          <button
            onclick={() => currPage = page}
            class="pagination-link"
            class:is-current={page === currPage}
            aria-label="Goto page {page}">{page}</button
          >
        </li>
      {/each}

    </ul>
  </nav>
</div>
