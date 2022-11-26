def isbalanced(s):     
    s_=""     
    while(s!=s_):         
        s_=s         
        s=((s.replace('()','')).replace('[]','')).replace('{}','')     
        if(s==""):         
            print("valid")      
        else:         
            print("invalid")  
            #Reading data from 
            file f = open('data.txt','r')  
            #Reading each line as list of elements 
            s = f.readlines()  
            #print(s) 
            allowed=['[',']','{','}','(',')']  
            #removing all the unneccessary characters  
            for i in range(len(s)):     
                for j in s[i]:         
                    if j not in allowed:             
                        s[i]=s[i].replace(j,"")  
                        #Checking only One or two Parentheses on each line for i in s:     
                        if len(i)>2:        
                            print("invalid")         
                            exit(1)  

    isbalanced("".join(s))
