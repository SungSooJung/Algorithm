t=[4,1,4,7,6]

def solution(price):
    answer = []
    for i in range(0,len(price)):
        day=0
        for j in range(i,len(price)):

            if price[i]<price[j]:
                answer.append(day)
                break
            elif j==len(price)-1:
                answer.append(-1)
            else:
                day+=1
        
    return answer


print(solution(t))