<script lang="ts">
  import type {ComponentType} from 'svelte';
  type Props = {
    value: string,
    type: string, 
    placeholder: string,
    icon: ComponentType,
  };

  let {
    value = $bindable(),
    type = "text",
    placeholder = "",
    icon,
  }: Props = $props();
</script>

<div class="field">
  <div class="control has-icons-left">
    <input bind:value={value} class="input" {type} {placeholder} />
    <span class="icon is-small is-left">
      <svelte:component this={icon} />
    </span>
  </div>
</div>
