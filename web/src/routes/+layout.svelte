<script>
  import 'bulma/css/bulma.css';
  import NavBar from '$lib/components/NavBar.svelte';
  import Footer from '$lib/components/Footer.svelte';
  import { createAuthStore } from '$lib/stores/auth.svelte';

  const {children} = $props();

  const authStore = createAuthStore();
  $effect.pre(() => authStore.checkSessionValidity());
</script>

<NavBar />
{@render children()}
<Footer />
