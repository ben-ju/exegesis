<script setup>
import phrasingImg from "../../assets/phrasing.png"
</script>

<template>
  <section class="bg-white py-12">
    <div class="container mx-auto px-4">
      <div class="flex flex-col md:flex-row items-center">
        <div class="md:w-1/2">
          <h1 class="text-4xl font-bold text-gray-800">
            Découvrez notre outil d'étude biblique
          </h1>
          <p class="mt-4 text-gray-600">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer nec odio. Praesent libero. Sed cursus ante
            dapibus diam.
          </p>
          <div class="mt-6">
            <a href="#"
              class="bg-blue-600 text-white py-3 px-6 rounded shadow hover:bg-blue-700 transition duration-300">
              Commencer
            </a>
          </div>
        </div>
        <div class="md:w-1/2 mt-8 md:mt-0">
          <img :src="phrasingImg" alt="Hero Image" class="w-full rounded shadow" />
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
</style>
