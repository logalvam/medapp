<template>
    <div>
        <v-container>
            <v-card>
                <v-data-table
                :headers="headers"
                :items="history"
                :items-per-page="5"
                class="elevation-1"
            ></v-data-table>
            <!-- {{ history }} -->
            </v-card>
        </v-container>
    </div>
</template>
<script>
  export default {
    data () {
      return {
        history:this.$store.state.loginhistory,
        headers: [
          {
            text: 'User ID',
            align: 'start',
            sortable: false,
            value: 'userid',
          },
          { text: 'LogIn', value: 'login' },
          { text: 'Log Out', value: 'logout' },
        ],
      }
    },
  }
</script>