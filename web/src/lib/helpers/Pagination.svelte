<script lang="ts">
  // one should bind to currPage to know what page to render
  type Props = { currPage: number; maxPage: number };
  let { currPage = $bindable(0), maxPage = $bindable(10) }: Props = $props();

  function range(begin: number, end: number) {
    let pages = [];
    for (let i = begin; i <= end; i++) {
      pages.push(i);
    }
    return pages;
  }

  function prevPage() {
    if (currPage > 0) --currPage;
  }

  function nextPage() {
    if (currPage < maxPage - 1) ++currPage;
  }

  let pages = $derived.by(() => {
    let minShownPage = currPage - 4;
    if (minShownPage < 0) minShownPage = 0;
    let maxShownPage = currPage + 5;
    if (maxShownPage > maxPage) maxShownPage = maxPage;
    return range(minShownPage, maxShownPage);
  });
</script>

<div class="is-flex is-justify-content-center">
  <nav class="pagination" aria-label="pagination">
    <button onclick={prevPage} class="pagination-previous">Prev</button>
    <button onclick={nextPage} class="pagination-next">Next</button>

    <ul class="pagination-list">
      {#if currPage > maxPage}
        <li>
          <button class="pagination-link" aria-label="Goto page 1">1</button>
        </li>
        <li>
          <span class="pagination-ellipsis">&hellip;</span>
        </li>
      {/if}

      {#each pages as page}
        <li>
          <button
            onclick={() => (currPage = page)}
            class="pagination-link"
            class:is-current={page === currPage}
            aria-label="Goto page {page}">{page + 1}</button
          >
        </li>
      {/each}
    </ul>
  </nav>
</div>
