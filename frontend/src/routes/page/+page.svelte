<script>
    import { name } from '$lib/stores/name';
    import { goto } from '$app/navigation';
</script>

<h1>Hello {$name}!</h1>
<button on:click={() => goto('/')}>/</button>