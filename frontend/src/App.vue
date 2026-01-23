<script setup>
import { ref } from 'vue'
import { GenerateTicket } from '../wailsjs/go/wailsapp/App'

const numbers = ref([])
const error = ref('')

function generate() {
  GenerateTicket().then(result => {
    numbers.value = result.Numbers
    error.value = ''
  }).catch(err => {
    error.value = err
  })
}
</script>

<template>
  <v-app>
    <v-main>
      <v-container class="fill-height justify-center">
        <v-card class="elevation-12 rounded-xl" min-width="400" max-width="600">
          <v-card-title class="text-h4 text-center py-6 font-weight-bold text-primary">
            6/49 Lottery
          </v-card-title>

          <v-card-text class="text-center py-8">
            <v-row justify="center" v-if="numbers.length > 0">
              <v-col v-for="num in numbers" :key="num" cols="auto">
                <v-avatar color="primary" size="50" class="text-h5 font-weight-bold">
                  {{ num }}
                </v-avatar>
              </v-col>
            </v-row>
            <div v-else class="text-h6 text-medium-emphasis">
              Click generate to get your lucky numbers!
            </div>
            
            <v-alert v-if="error" type="error" class="mt-4" variant="tonal">
              {{ error }}
            </v-alert>
          </v-card-text>

          <v-card-actions class="justify-center pb-6">
            <v-btn
              color="secondary"
              size="x-large"
              variant="flat"
              prepend-icon="mdi-ticket"
              @click="generate"
              class="rounded-pill px-8"
            >
              Generate Ticket
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-container>
    </v-main>
  </v-app>
</template>

<style>
/* Vuetify handles most styles, but we can add custom overrides here if needed */
</style>
