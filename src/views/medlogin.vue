<template>
    <div>
        <v-app-bar
            color="blue-grey lighten-1"
            elevation="4"
            >
          <v-container fluid class="black">
            <v-row class="d-flex flex-row pa-1 justify-end">
                <!-- <v-col class=""> -->
                    <v-col cols="12" >
                        <v-btn class="ml-3">Dashboard</v-btn>
                        <v-btn  @click="bill" class="ml-3">Bill Entry</v-btn>
                        <v-btn @click="stockview" class="ml-3">StockView</v-btn>
                        <v-btn  @click="stockentry" class="ml-3">stock Entry</v-btn>
                        <v-btn  class="ml-3" @click="report">Sales Report</v-btn>
                        <v-btn class="ml-3" @click="adduser">Adduser</v-btn>
                        <v-btn  class="ml-3" @click="loginhis">login history</v-btn>
                        <v-btn class="ml-3" @click="logout">login out</v-btn>
                    </v-col>


                <!-- </v-col> -->
            </v-row>
            </v-container>
        </v-app-bar>
        <!-- Adduser -->
        <v-card>
            <stockview v-show="viewshow"/>
        </v-card>
        <v-card>
            <bill v-show="billshow"/>
        </v-card>
        <v-card >
            <stockentry v-show="entrystock" />
        </v-card>
        <v-card>
            <adduser v-show="useradd"/>
        </v-card>
        <v-card>
            <loginhistory v-show="loginhistory"/>
        </v-card>
        <v-card>
            <salesreport v-show="salesreport"/>
        </v-card>

      
    </div>

</template>
<script>
import stockview from "../components/stockview.vue"
import bill from "../components/bill.vue"
import stockentry from "../components/stockbar.vue"
import adduser from "../components/useraddbar.vue"
import loginhistory from "../components/loginhistory.vue"
import salesreport from "../components/salesreport.vue"
export default{
    name:'',
    data(){
        return{
            // addUser:true,
            // stockEntry:true,
            // stockView:true,
            // billEntry:true,
            // SaleReport:true,
            // loginHistory:true,
            biller:false,
            manager:false,
            admin:false,
            inventry:false,
            viewshow:false,
            billshow:false,
            entrystock:false,
            useradd:false,
            loginhistory:false,
            salesreport:false,
            logoudate:'',
            medicine:this.$store.state.medicinemaster,
            stock:this.$store.state.stock,  
            newarr:[]
        }
    },
    props:{
    
    },
    components:{
        stockview,
        bill,
        stockentry,
        adduser,
        loginhistory,
        salesreport
    },
    methods:{
        logout(){
            this.logoudate = new Date().toLocaleString()
            this.$router.push("/")
            this.$emit('logout',this.logoudate)
            
        },
        adduser(){
            this.useradd=true
            this.billshow=false
            this.viewshow=false
            this.entrystock=false
            this.loginhistory=false
            this.salesreport=false

            // this.$router.push('/adduser')
        },
        bill(){
            this.billshow=true
            this.viewshow=false
            this.entrystock=false
            this.useradd=false
            this.loginhistory=false
            this.salesreport=false
            


            // this.$router.push('/bill')
        },
        stockview() {
            this.viewshow=true
            this.billshow=false
            this.entrystock=false
            this.useradd=false
            this.loginhistory=false
            this.salesreport=false
            // this.$emit('stockarr')
            // this.$router.push("/stockView1");

        },
        stockentry() {
            this.entrystock=true
            this.viewshow=false
            this.billshow=false
            this.useradd=false
            this.loginhistory=false
            this.salesreport=false

            // this.$router.push("/stockentry")
        },
        loginhis(){
            this.loginhistory=true
            this.entrystock=false
            this.viewshow=false
            this.billshow=false
            this.useradd=false
            this.salesreport=false

        },
        report(){
            this.salesreport=true
            this.loginhistory=false
            this.entrystock=false
            this.viewshow=false
            this.billshow=false
            this.useradd=false
            // this.$router.push('/salesreport')
        },
        
    },
    computed:{

    },
    watch:{
        
    }
}
</script>