<template>
  <v-flex xs12>
    <v-card>
      <v-card-title primary-title>
        <div>
          <h3 class="headline mb-0">Drip Rate - {{ dripsPerMinute }}</h3>
        </div>
      </v-card-title>
      <v-card-text>
        <v-slider min="0" max="120" v-model="dripsPerMinute" step="0"></v-slider>
        <v-btn color="primary" v-on:click.native="startDripper">On</v-btn>
        <v-btn color="primary" v-on:click.native="stopDripper">Off</v-btn>
        <v-btn color="primary" v-on:click.native="runDripper">Run</v-btn>
      </v-card-text>
    </v-card>
  </v-flex>
</template>

<script>
export default {
  data () {
    return {
      dripsPerMinute: 0,
      dripperState: 'off'
    }
  },

  beforeMount () {
    this.getDripperState()
  },

  beforeUpdate () {
    if (this.dripperState === 'drip') {
      this.startDripper()
    }
  },

  methods: {
    getDripperState () {
      this.$http.get('/api/cold-brew/v1/dripper')
        .then(response => {
          this.dripsPerMinute = response.body.dripsPerMinute
          this.dripperState = response.body.state
        }, error => {
          console.error(error)
        })
    },

    startDripper () {
      this.$http.post('/api/cold-brew/v1/dripper/drip', {dripsPerMinute: this.dripsPerMinute})
        .then(response => {
          this.dripperState = response.body.state
        }, error => {
          console.error(error)
        })
    },

    stopDripper () {
      this.$http.post('/api/cold-brew/v1/dripper/off')
        .then(response => {
          this.dripperState = response.body.state
        }, error => {
          console.error(error)
        })
    },

    runDripper () {
      this.$http.post('/api/cold-brew/v1/dripper/run')
        .then(response => {
          this.dripperState = response.body.state
        }, error => {
          console.error(error)
        })
    }
  },

  name: 'ManualControls'
}
</script>