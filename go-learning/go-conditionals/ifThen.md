<h2 align="center">
As the name indicates it is quite simple. If a certain condition is satisfied, then perform some action.
</h2>


***Structure of If Then Statement
+---------+
|  Start  | 
+---------+
     |
	 V
	+-------------------+
	| Condition			|-----------+
	+-------------------+			|
		|							|
		| True						| False					
		| 							|
		V							|
	+-----------------+				|
	|  Statements	  |				|
	+-----------------+				|
		|							|
		|							|
		V							|
+---------+							|
|  End	  |-------------------------+ 
+---------+


In the above example, we are checking if the accountBalance is less than 1000. If the accountBalance is less than 1000 we are printing a message to close the Account. We initialized the accountBalance to zero, so the output displays that we closed the Account. In a real world scenario this can be the case to check an account has minimum balance.