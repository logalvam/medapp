<template>
    <div>

    <v-card width="100vw" height="70">
        <v-row>
      <v-col
        cols="12"
        sm="6"
        md="4"
      >
        <v-menu
          ref="menu"
          v-model="menu"
          :close-on-content-click="false"
          :return-value.sync="date1"
          transition="scale-transition"
          offset-y
          min-width="auto"
        >
          <template v-slot:activator="{ on, attrs }">
            <v-text-field
              v-model="date1"
              label="Start date"
              prepend-icon="mdi-calendar"
              readonly
              v-bind="attrs"
              v-on="on"
            ></v-text-field>
          </template>
          <v-date-picker
            v-model="date1"
            no-title
            scrollable
          >
            <v-spacer></v-spacer>
            <v-btn
              text
              color="primary"
              @click="menu = false"
            >
              Cancel
            </v-btn>
            <v-btn
              text
              color="primary"
              @click="$refs.menu.save(date1)"
            >
              OK
            </v-btn>
          </v-date-picker>
        </v-menu>
      </v-col>
      <v-spacer></v-spacer>
      <v-col
        cols="12"
        sm="6"
        md="4"
      >
        <v-dialog
          ref="dialog"
          v-model="modal"
          :return-value.sync="date2"
          persistent
          width="290px"
        >
          <template v-slot:activator="{ on, attrs }">
            <v-text-field
              v-model="date2"
              label="End date"
              prepend-icon="mdi-calendar"
              readonly
              v-bind="attrs"
              v-on="on"
            ></v-text-field>
          </template>
          <v-date-picker
            v-model="date2"
            scrollable
          >
            <v-spacer></v-spacer>
            <v-btn
              text
              color="primary"
              @click="modal = false"
            >
              Cancel
            </v-btn>
            <v-btn
              text
              color="primary"
              @click="$refs.dialog.save(date2)"
            >
              OK
            </v-btn>
          </v-date-picker>
        </v-dialog>
      </v-col>
      <v-col
        cols="12"
        sm="6"
        md="4"
      >
        <v-btn color='primary' @click="search">Search</v-btn>
      </v-col>
      <v-spacer></v-spacer>
      
    </v-row>
    </v-card>
    <template>
  <v-data-table
    :headers="headers"
    :items="filterarr"
    :items-per-page="5"
    class="elevation-1"
  ></v-data-table>
</template>
{{ arr }}
</div>

</template>
<script>
  export default {
    data () {
      return {
        date2: (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10),
    date1: (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10),
    menu: false,
    modal: false,
    arr:[],
    billdetails:this.$store.state.billdetails,
    billmaster:this.$store.state.billmaster,
    headers: [
          {
            text: 'Billno)',
            align: 'start',
            sortable: false,
            value: 'billno',
          },
          { text: 'Billdate', value: 'billdate' },
          { text: 'Medicine Name', value: 'medicinename' },
          { text: 'Quantity', value: 'quantity' },
          { text: 'unitprice', value: 'unitprice' },
          { text: 'amount', value: 'amount' },
        ],
        newarr:[],
        filterarr:[]
      }
    },
    mounted(){
        // console.log(this.billdetails)
        // console.log('kil')
        // console.log(this.billdetails)
        // console.log(this.billmaster)
        // newitemadd(){
        for(var i in this.billdetails){
            for(var j in this.billmaster){
                // console.log('ll')
                // console.log(this.billmaster)
                if(this.billdetails[i].billno === this.billmaster[j].billno){
                    // console.log('sat')
                    let mergedObject = {
                        billno: this.billdetails[i].billno,
                        medicinename: this.billdetails[i].medicinename, 
                        quantity: this.billdetails[i].qty, 
                        unitprice: this.billdetails[i].unitprice,
                        amount: Number(this.billdetails[i].qty)*Number(this.billdetails[i].unitprice), 
                        billdate: this.billmaster[j].billdate
                    }
                    this.newarr.push(mergedObject)
                }
            }
              }  
            //     // return newarr
        // }
        console.log(this.newarr)
    },
    computed:{

    },
    props:{
        generate:Boolean
    },
   methods:{
    search(){
      this.filterarr=[]
        let start = new Date( this.date).toLocaleDateString()
        let stop = new Date( this.date1).toLocaleDateString()
        for (var i in this.newarr){
            if(start>=this.newarr[i].billdate){
                if(stop<=this.newarr[i].billdate){
                    console.log('work')
                    console.log(this.newarr[i])
                    this.filterarr.push(this.newarr[i])
                }
            }
        }  
        console.log(this.filterarr)      

    }
   },
   watch:{
   },
  }
</script>