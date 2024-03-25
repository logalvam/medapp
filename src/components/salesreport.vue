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
              :min='today'
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
        <v-btn class="black white--text" @click="search">Search</v-btn>
      </v-col>
      <v-spacer></v-spacer>
      
    </v-row>
    </v-card>
    <template>
  <v-data-table
    :headers="headers"
    :items="report"
    :items-per-page="5"
    class="elevation-1"
    v-show="visible"
  ></v-data-table>
</template>
</div>

</template>
<script>
  export default {
    data () {
      return {
        date1: (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10),
    date2: (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10),
    menu: false,
    modal: false,
    arr:[],
    billdetails:this.$store.state.billdetails,
    billmaster:this.$store.state.billmaster,
    headers: [
          {
            text: 'Billno',
            align: 'start',
            sortable: false,
            value: 'billno',
          },
          { text: 'Billdate', value: 'billdate' },
          { text: 'Medicine Name', value: 'medicinename' },
          { text: 'Quantity', value: 'quantity' },
          // { text: 'unitprice', value: 'unitprice' },
          { text: 'amount', value: 'amount' },
        ],
        filterarr:[],
        visible:false,
        billno:0,
        billdate:'',
        qty:0,
        medicinename:'',
        amount:0,
        report:[],
      }
    },
    props:{
        generate:Boolean
    },
   methods:{
    search(){
      let newarray=[]
        let start = new Date( this.date).toLocaleDateString()
        let stop = new Date( this.date1).toLocaleDateString()
        for (var i in this.filterarr){
          // console.log(this.filterarr[i].billdate)
          if(start>=this.filterarr[i].billdate){
            if(stop<=this.filterarr[i].billdate){
                      newarray.push(this.filterarr[i])
                  }
              }
            }
            this.report=newarray 
    },
    
  },
   watch:{
   },
   mounted(){
    // console.log('  mounted')
    
this.visible=true
      let mergedArray = [];

    for (let i = 0; i < this.billdetails.length; i++) {
    const current = this.billdetails[i];
    const match = this.billmaster.find(item => item.billno === current.billno);
    if (match) {
        mergedArray.push({
            billno: current.billno,
            billdate: match.billdate,
            medicinename: current.medicinename,
            quantity: current.quantity,
            amount: current.unitprice
        });
    }
    this.filterarr = mergedArray
} 
          // console.log(this.filterarr)

             
   },
   beforeMount(){
    // console.log('before mounted')
    this.today=new Date().toLocaleDateString()
    // console.log(this.today)
    },
    
  }
</script>

