<script>
  import { Router, Route } from 'svelte-routing'
  import Home from './pages/Home.svelte'
  import About from './pages/About.svelte'
  import NavLink from './components/NavLink.svelte'
  export let url = ''
</script>

<Router {url} class="">
  <nav class="p-3 bg-blue-500 text-white select-none">
    <NavLink to="/">Home</NavLink>
    <NavLink to="about">About</NavLink>
  </nav>
  <div>
    <Route path="about" component={About} />
    <Route path="/"><Home /></Route>
  </div>
</Router>
