<script lang="ts">
  type Props = {
    value: string;
    items: string[];
    size?: "is-small" | "" | "is-medium" | "is-large";
  };

  let {
    items = ["1-10", "10-20", "20-30"],
    value = $bindable(items[0]),
    size = "",
  }: Props = $props();

  let selectedItem = $state(0);

  function chooseItem(i: number) {
    selectedItem = i;
    value = items[i];
  }

  // choose first in list by default
  chooseItem(0);

  let isOpen = $state(false);
  function toggleDropdown() {
    isOpen = !isOpen;
  }
</script>

<div class="dropdown" class:is-active={isOpen} onclick={toggleDropdown}>
  <div class="dropdown-trigger">
    <button
      class="button {size}"
      aria-haspopup="true"
      aria-controls="dropdown-menu"
    >
      <span>{value}</span>
      <span class="icon is-small">
      </span>
    </button>
  </div>
  <div class="dropdown-menu" id="dropdown-menu" role="menu">
    <div class="dropdown-content">
      {#each items as item, i (i)}
        <a
          class="dropdown-item"
          class:is-active={i === selectedItem}
          onclick={() => chooseItem(i)}>{item}</a
        >
      {/each}
    </div>
  </div>
</div>
