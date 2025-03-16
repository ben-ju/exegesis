<script setup>

</script>
<template>
  <nav class="bg-blue-600 shadow">
    <div class="max-w-7xl mx-auto px-4">
      <div class="flex justify-between">
        <!-- Section gauche : Logo et navigation principale -->
        <div class="flex space-x-4">
          <!-- Logo -->
          <div>
            <router-link to="/" class="flex items-center py-5 px-2 text-white">
              <!-- Exemple de logo SVG, remplacez-le par votre logo si nécessaire -->
              <svg class="h-6 w-6 mr-1" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6l4 2" />
              </svg>
              <span class="font-bold">MaMarque</span>
            </router-link>
          </div>
          <!-- Liens de navigation principaux (affichés à partir de md) -->
          <div class="hidden md:flex items-center space-x-1">
            <router-link to="/" class="py-5 px-3 text-white hover:text-blue-200">Accueil</router-link>
            <router-link to="/about" class="py-5 px-3 text-white hover:text-blue-200">À Propos</router-link>
            <router-link to="/services" class="py-5 px-3 text-white hover:text-blue-200">Services</router-link>
            <router-link to="/contact" class="py-5 px-3 text-white hover:text-blue-200">Contact</router-link>
          </div>
        </div>
        <!-- Section droite : Actions secondaires -->
        <div class="hidden md:flex items-center space-x-1">
          <router-link to="/login" class="py-5 px-3 text-white hover:text-blue-200">Connexion</router-link>
          <router-link to="/register"
            class="py-2 px-3 bg-blue-500 hover:bg-blue-400 text-white rounded transition duration-300">
            Inscription
          </router-link>
        </div>
        <!-- Bouton mobile -->
        <div class="md:hidden flex items-center">
          <button @click="toggleMobileMenu" class="mobile-menu-button focus:outline-none">
            <svg class="w-6 h-6 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
              stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
          </button>
        </div>
      </div>
    </div>
    <!-- Menu mobile -->
    <div v-if="mobileMenuOpen" class="md:hidden">
      <router-link to="/" class="block py-2 px-4 text-sm text-white hover:bg-blue-700">Accueil</router-link>
      <router-link to="/about" class="block py-2 px-4 text-sm text-white hover:bg-blue-700">À Propos</router-link>
      <router-link to="/services" class="block py-2 px-4 text-sm text-white hover:bg-blue-700">Services</router-link>
      <router-link to="/contact" class="block py-2 px-4 text-sm text-white hover:bg-blue-700">Contact</router-link>
      <router-link to="/login" class="block py-2 px-4 text-sm text-white hover:bg-blue-700">Connexion</router-link>
      <router-link to="/register" class="block py-2 px-4 text-sm text-white hover:bg-blue-700">Inscription</router-link>
    </div>
  </nav>
</template>
