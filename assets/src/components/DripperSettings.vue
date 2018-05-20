<template>
  <v-flex xs12>
    <v-card>
      <v-card-title primary-title>
        <div>
          <h3 class="headline mb-0">Settings</h3>
        </div>
      </v-card-title>
      <v-card-text>
        <p>Drip Duration - {{ dripDuration }}</p>
        <v-slider min="150" max="300" v-model="dripDuration" step="1"></v-slider>
        <p>Drip Speed - {{ dripSpeed }}</p>
        <v-slider min="0" max="120" v-model="dripSpeed" step="1"></v-slider>
        <p>Run Speed - {{ runSpeed }}</p>
        <v-slider min="0" max="255" v-model="runSpeed" step="1"></v-slider>
        <v-btn color="primary" v-on:click.native="setDripperSettings">Submit</v-btn>
      </v-card-text>
    </v-card>
  </v-flex>
</template>

<script>
export default {
  data () {
    return {
      dripDuration: 250,
      dripSpeed: 100,
      runSpeed: 255
    }
  },

  beforeMount () {
    this.getDripperSettings()
  },

  methods: {
    getDripperSettings () {
      this.$http.get('/api/cold-brew/v1/dripper/settings')
        .then(response => {
          this.dripDuration = response.body.dripDuration
          this.dripSpeed = response.body.dripSpeed
          this.runSpeed = response.body.runSpeed
        }, error => {
          console.error(error)
        })
    },

    setDripperSettings () {
      this.$http.post('/api/cold-brew/v1/dripper/settings', {
        dripDuration: this.dripDuration,
        dripSpeed: this.dripSpeed,
        runSpeed: this.runSpeed
      })
        .then(response => {
          this.dripDuration = response.body.dripDuration
          this.dripSpeed = response.body.dripSpeed
          this.runSpeed = response.body.runSpeed
        }, error => {
          console.error(error)
        })
    }
  },

  name: 'DripperSettings'
}
</script>
