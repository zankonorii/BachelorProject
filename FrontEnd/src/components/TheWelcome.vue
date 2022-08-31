<script setup>
import WelcomeItem from './WelcomeItem.vue'
import DocumentationIcon from './icons/IconDocumentation.vue'
import ToolingIcon from './icons/IconTooling.vue'
import EcosystemIcon from './icons/IconEcosystem.vue'
import CommunityIcon from './icons/IconCommunity.vue'
import SupportIcon from './icons/IconSupport.vue'
</script>

<template>
  <WelcomeItem>
    <template #icon>
      <DocumentationIcon />
    </template>
    <template #heading>Documentation</template>

    Full document and source code on <a target="_blank" href="https://github.com/zankonorii/BachelorProject">github</a> .
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <ToolingIcon />
    </template>
    <template #heading>Tooling</template>

    This project is served and bundled with
    <a href="https://vitejs.dev/guide/features.html" target="_blank">Vite</a>. The recommended IDE
    setup is <a href="https://code.visualstudio.com/" target="_blank">VSCode</a> +
    <a href="https://github.com/johnsoncodehk/volar" target="_blank">Volar</a>. If you need to test
    your components and web pages, check out
    <a href="https://www.cypress.io/" target="_blank">Cypress</a> and
    <a href="https://on.cypress.io/component" target="_blank"
      >Cypress Component Testing</a
    >.

    <br />

    More instructions are available in <code>README.md</code>.
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <CommunityIcon />
    </template>
    <template #heading>LogIn & Registration</template>
    You Can <RouterLink to="/login">LogIn</RouterLink>  or  <RouterLink to="/register">Register</RouterLink>  in this website.
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <SupportIcon />
    </template>
    <template #heading>About Me</template>

    I am Zanko Nori, software engineer student at  <a target="_blank" href="http://www.azaruniv.ac.ir/">shahid madani</a>. univercist.
   
  </WelcomeItem>
</template>
