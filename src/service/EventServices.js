import axios from "axios";
const baseApiClient = axios.create({

  baseURL: `http://localhost:8260`,
  withCredentials: false,
  headers: {
    Accept: `application/json`,
    "Content-Type": `application/json`,
  },
});

export default{
    LoginValidation(body){
        return baseApiClient.put('/loginValidation',body);
    },
    loginHistory(user){
      console.log(user)
      const hdr ={
          headers:{"USER": user}
      }
      return baseApiClient.put('/loginhistoryInsert',hdr);
    }
    
}