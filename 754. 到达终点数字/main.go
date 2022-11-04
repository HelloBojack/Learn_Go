package main
func reachNumber(target int) int {
	if target<0{
			target=-target
	}
	sum:=0
	for i:=1;;i++ {
			sum+=i
			if sum>=target {
					d:=sum-target
					if d%2==0{
						return i
					}
							if (sum+i-target)%2==1{
								return i+1
							}else{
								return i+2
							}
			}
	}
}