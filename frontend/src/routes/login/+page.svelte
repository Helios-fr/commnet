<script>
    import { Input } from '$lib/components/ui/input/index.js';
    import { Button } from '$lib/components/ui/button/index.js';
    import { Label } from '$lib/components/ui/label/index.js';

    import { username, isLoggedIn } from '$lib/stores/username';
    import { goto } from '$app/navigation';

    let name

    async function handleSubmit() {
        console.log('handleSubmit')
        // set the button to loading state
        let button = document.querySelector('button')
        let input = document.querySelector('input')
        button.innerHTML = 'Loading...'
        input.disabled = true
        button.disabled = true
        name = await username.login(name)
        if (name == false) {
            button.innerHTML = 'Start chatting'
            button.disabled = false
            input.disabled = false
            return
        }
        $isLoggedIn = true
        goto('/')
    }
</script>

<h1>Welcome back to Commnet 👋</h1>

<!-- Do this with shadcn form -->
<form on:submit|preventDefault={handleSubmit}>
    <Label for="name">Enter your name to login:</Label>
    <Input bind:value={name} id="name" placeholder="Your name" class="my-2" />
    <Button disabled={!name}>Login</Button>
    <Button on:click={() => goto('/')}>/</Button>
</form>