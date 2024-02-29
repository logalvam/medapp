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
                        <v-btn  class="ml-3" @click="dashboard">Dashboard</v-btn>
                        <v-btn v-show="billEntry"   @click="bill" class="ml-3">Bill Entry</v-btn>
                        <v-btn v-show="stockView"  @click="stockview" class="ml-3">StockView</v-btn>
                        <v-btn v-show="stockEntry"   @click="stockentry" class="ml-3">stock Entry</v-btn>
                        <v-btn v-show="salesReport"   class="ml-3" @click="report">Sales Report</v-btn>
                        <v-btn v-show="addUser"  class="ml-3" @click="adduser">Adduser</v-btn>
                        <v-btn v-show="loginHistory"   class="ml-3" @click="loginhis">login history</v-btn>
                        <v-btn   class="ml-3" @click="logout">log out</v-btn>
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
            <bill :currentuserid="currentuserid" v-show="billshow"/>
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
            <salesreport :generate="generate" v-show="salesreport"/>
        </v-card>
        <v-card>
            <billerdash :val="currentrole" v-show="billdash" />
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
import billerdash from "../components/billerdash.vue"
export default{
    name:'medlogin',
    data(){
        return{
            addUser:false,
            stockEntry:false,
            stockView:false,
            billEntry:false,
            salesReport:false,
            loginHistory:false,


            biller:false,
            manager:false,
            admin:false,
            inventry:false,
            viewshow:false,
            billshow:false,
            entrystock:false,
            useradd:false,
            billdash:false,
            loginhistory:false,
            salesreport:false,
            logoudate:'',
            // value:'',

            medicine:this.$store.state.medicinemaster,
            stock:this.$store.state.stock,  
            login:this.$store.state.login,
            newarr:[],
            loghistory:this.$store.state.loginhistory,
            currentuserid:'',
            currentrole:'',
            generate:false
        }
    },
    props:{
        userrole:String,
        Euserid:String    
    },
    components:{
        stockview,
        bill,
        stockentry,
        adduser,
        loginhistory,
        salesreport,
        billerdash
    },
    methods:{
        logout(){
            this.logoudate = new Date().toLocaleString()
            for (var i in this.loghistory){
                if(this.currentuserid ===  this.loghistory[i].userid){
                    this.loghistory[i].logout = this.logoudate
                }
            }
            this.$router.push("/")
            
        },
        adduser(){
            this.useradd=true
            this.billshow=false
            this.viewshow=false
            this.entrystock=false
            this.loginhistory=false
            this.salesreport=false
            this.billdash=false

            // this.$router.push('/adduser')
        },
        bill(){
            this.billshow=true
            this.viewshow=false
            this.entrystock=false
            this.useradd=false
            this.loginhistory=false
            this.salesreport=false
            this.billdash=false


            // this.$router.push('/bill')
        },
        dashboard(){
            this.billdash=true
            this.billshow=false
            this.viewshow=false
            this.entrystock=false
            this.useradd=false
            this.loginhistory=false
            this.salesreport=false
            this.value='total'
            // if(this.currentrole==='Biller'){

            // }
        }
        ,
        stockview() {
            this.viewshow=true
            this.billshow=false
            this.entrystock=false
            this.useradd=false
            this.loginhistory=false
            this.salesreport=false
            // this.$emit('stockarr')
            // this.$router.push("/stockView1");
            this.billdash=false

        },
        stockentry() {
            this.entrystock=true
            this.viewshow=false
            this.billshow=false
            this.useradd=false
            this.loginhistory=false
            this.salesreport=false
            this.billdash=false
            // this.$router.push("/stockentry")
        },
        loginhis(){
            this.loginhistory=true
            this.entrystock=false
            this.viewshow=false
            this.billshow=false
            this.useradd=false
            this.salesreport=false
            this.billdash=false

        },
        report(){
            this.salesreport=true
            this.loginhistory=false
            this.entrystock=false
            this.viewshow=false
            this.billshow=false
            this.useradd=false
            this.billdash=false
            this.generate=true
            setTimeout(() => {
                this.generate=false
            }, 2000);
            // this.$router.push('/salesreport')
        },
        
    },
    created(){
        
        let len = this.loghistory.length-1
            this.currentuserid=this.loghistory[len].userid
            for (var i in this.login){
                if(this.currentuserid=== this.login[i].username){
                    this.currentrole=this.login[i].role
                }

            }
    },
    watch:{
    currentrole:{
        handler(){
            if(this.currentrole==='Biller'){
                this.stockView=true
                this.billEntry=true
            }
            else if(this.currentrole === 'Manager'){
                this.stockEntry=true
                this.stockView=true
                this.salesReport=true
            }
            else if(this.currentrole === 'SystemAdmin'){
                this.addUser=true
                this.loginHistory=true
            }
            else if(this.currentrole === 'Inventry'){
                this.stockEntry=true
                this.stockView=true
            }
        },immediate:true
    }
}
    
}
</script>

