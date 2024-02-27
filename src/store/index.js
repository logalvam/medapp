import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    login:[
      {username:'abcd',password:'1234',role:'Biller'},
      {username:'efgh',password:'1234',role:'Biller'},
      {username:'manager',password:'1234',role:'Manager'},
      {username:'invent',password:'5678',role:'Inventry'},
      {username:'admin',password:'9012',role:'SystemAdmin'},
    ],
    loginhistory:[
    ],
    medicinemaster:[
      {medicinename:'Bacitracin',brandname:'xyz'},
      {medicinename:'Bactrim',brandname:'xyz'},
      {medicinename:'Aspirin',brandname:'Bayer Aspirin'},
      {medicinename:'Levothyroxine',brandname:'Levoxyl'},
      {medicinename:'Metoprolol',brandname:'Lopressor'},
    ],
    stock:[
      {medicinename:'Metoprolol',quantity:5,amount:10},
      {medicinename:'Bacitracin',quantity:3,amount:12},
      {medicinename:'Aspirin',quantity:10,amount:20},
      {medicinename:'Levothyroxine',quantity:15,amount:4},
      {medicinename:'Bactrim',quantity:6,amount:13},

    ],
    billmaster:[

    ],
    billdetails:[

    ]
  },
  mutations: {},
  actions: {},
  modules: {},
});
