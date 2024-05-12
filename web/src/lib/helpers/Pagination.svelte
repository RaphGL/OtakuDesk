<!-- todo: improve how pages are generated -->
<script>
  // one should bind to currPage to know what page to render
  let { currPage = $bindable(1), maxPage  = 10} = $props();

  let pages = [];
  for (let i = currPage; i <= maxPage; i++) {
    pages.push(i);
  }
</script>

<div class="is-flex is-justify-content-center">
  <nav class="pagination" role="pagination" aria-label="pagination">
    <a href="" class="pagination-previous">Prev</a>
    <a href="" class="pagination-next">Next</a>

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
