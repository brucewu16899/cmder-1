GOOF----LE-8-2.0CA      ]  4  h      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  srfi¤	g  srfi-18¤	 ¤		g  filenameS¤	
f  srfi/srfi-18.scm¤	g  importsS¤	g  srfi-34¤	 ¤	 ¤	 ¤	g  exportsS¤	g  make-thread¤	g  thread-name¤	g  thread-specific¤	g  thread-specific-set!¤	g  thread-start!¤	g  thread-yield!¤	g  thread-sleep!¤	g  thread-terminate!¤	g  thread-join!¤	g  
make-mutex¤	g  
mutex-name¤	g  mutex-specific¤	g  mutex-specific-set!¤	g  mutex-state¤	g  mutex-lock!¤	 g  mutex-unlock!¤	!g  make-condition-variable¤	"g  condition-variable-name¤	#g  condition-variable-specific¤	$g   condition-variable-specific-set!¤	%g  condition-variable-signal!¤	&g  condition-variable-broadcast!¤	'g  condition-variable-wait!¤	(g  current-time¤	)g  time?¤	*g  time->seconds¤	+g  seconds->time¤	,g  current-exception-handler¤	-g  with-exception-handler¤	.g  raise¤	/g  join-timeout-exception?¤	0g  abandoned-mutex-exception?¤	1g  terminated-thread-exception?¤	2g  uncaught-exception?¤	3g  uncaught-exception-reason¤	4 !"#$%&'()*+,-./0123 #¤	5g  
re-exportsS¤	6g  current-thread¤	7g  thread?¤	8g  mutex?¤	9g  condition-variable?¤	:6789 ¤	;g  replacementsS¤	<(!. ¤	=g  set-current-module¤	>= ¤	?= ¤	@g  	provided?¤	Ag  threads¤	Bg  error¤	Cf  +SRFI-18 requires Guile with threads support¤	Dg  cond-expand-provide¤	Eg  current-module¤	F ¤	Gg  	scm-error¤	Hg  wrong-type-arg¤	If  Wrong type argument: ~S¤	Jg  check-arg-type¤	Kg  abandoned-mutex-exception¤	Lg  join-timeout-exception¤	Mg  terminated-thread-exception¤	Ng  uncaught-exception¤	Og  make-weak-key-hash-table¤	Pg  object-names¤	Qg  object-specifics¤	Rg  thread-start-conds¤	Sg  thread-exception-handlers¤	T. ¤	U. ¤	Vg  srfi-18-exception-preserver¤	Wg  initial-handler¤	Xg  make-object-property¤	Yg  thread->exception¤	Zg  setter¤	[Z ¤	\Z ¤	]g  srfi-18-exception-handler¤	^g  	hashq-ref¤	_g  
hashq-set!¤	`g  current-handler-stack¤	ag  
procedure?¤	bf  with-exception-handler¤	cg  thunk?¤	d- ¤	e- ¤	ff  uncaught-exception-reason¤	gg  launch-mutex¤	hg  launch-condition-variable¤	ig  start-mutex¤	jg  start-condition-variable¤	kg  
lock-mutex¤	lg  call-with-new-thread¤	mg  signal-condition-variable¤	ng  unlock-mutex¤	og  wait-condition-variable¤	pf  thread-name¤	qf  thread-specific¤	rf  thread-specific-set!¤	sf  thread-start!¤	tg  hashq-remove!¤	ug  yield¤	vg  number?¤	wf  thread-sleep!¤	xg  inexact->exact¤	yg  truncate¤	zg  sleep¤	{g  usleep¤	|g  wrap¤	}g  thread-cleanup¤	~g  set-thread-cleanup!¤	g  cancel-thread¤ g  join-thread¤ g  length¤  ¤  ¤ g  unchecked-unlock¤ g  allow-external-unlock¤ g  	recursive¤ f  
mutex-name¤ f  mutex-specific¤ f  mutex-specific-set!¤ g  mutex-owner¤ g  thread-exited?¤ g  	abandoned¤ g  mutex-level¤ g  	not-owned¤ g  not-abandoned¤ g  catch¤ g  abandoned-mutex-error¤ ! ¤ ! ¤ f  condition-variable-name¤ f  condition-variable-specific¤ f   condition-variable-specific-set!¤ g  broadcast-condition-variable¤ g  gettimeofday¤ g  integer?¤ f  time->seconds¤ f  seconds->time¤C 5   h80  É  ]4	
45:;<5 4? >  "  G   4@iA5$  "  4BiC>  "  G  4Di4Ei5 F>  "  G  GHI  h    Í   ]4 5$  C 6Å       g  pred
		  g  arg		  g  caller			   g  filenamef  srfi/srfi-18.scm
	`
		a			a			c			d			d	+		d	6	 	c	 			 	  g  nameg  check-arg-type CJRK KRL LRM MRN NR4Oi5 PR4Oi5 QR4Oi5 RR4Oi5 SRU.RVN   h      ] 6     x       g  obj
		  g  filenamef  srfi/srfi-18.scm
	s
			t			t	 		  g  nameg  initial-handler CWR4Xi5 YR12\Y6     h8   º   ]	4 5$  "  4 5$  4545  6C  ²       g  obj
		6 g  t			   g  filenamef  srfi/srfi-18.scm
	x
		y	
			y			z	
	$	y		,	{		4	{	 		6  g  nameg  srfi-18-exception-preserver CVRVN h0   Ã   - 1 3  &  C(  	 "   6  »       g  key
			. g  args			.  g  filenamef  srfi/srfi-18.scm
	}
	 		 		 	$	  	(	+ 	(	. 	 			.
  g  nameg  srfi-18-exception-handler C]R6^S_W    h0   º   ]45  4 5$  C  6      ²       g  ct
		* g  t		*  g  filenamef  srfi/srfi-18.scm
 
	 		 		
 		 		( 	1	* 	 		*
  g  nameg  current-handler-stack C`R6`Jabc_Se_S        h    h   ]4LL>  "  G  L  6  `       g  obj
		  g  filenamef  srfi/srfi-18.scm
 		 		 	 		   C_S        h0   q   ]4L>   G 4L L>  "  G   E      i       g  res
		*  g  filenamef  srfi/srfi-18.scm
 		 		 		* 	 		*
   C    hp   $  ]45 45 4 >  "  G  4>  "  G  4 >  "  G  	
 O O 6       g  handler
		o g  thunk		o g  ct			o g  hl			o  g  filenamef  srfi/srfi-18.scm
 
	 		 		 		 		 	'	 		' 		/ 	!	4 		= 		H 	-	M 		o 	 		o	  g  nameg  with-exception-handler C-R`   h   {   ] 45 C       s       g  filenamef  srfi/srfi-18.scm
 
	 		 	 			
  g  nameg  current-exception-handler C,RL    h      ] Cz       g  obj
		  g  filenamef  srfi/srfi-18.scm
  
	  	& 		  g  nameg  join-timeout-exception? C/RK     h      ] C}       g  obj
		  g  filenamef  srfi/srfi-18.scm
 ¡
	 ¡	) 		  g  nameg  abandoned-mutex-exception? C0RN  h      ] $   CC             g  obj
		  g  filenamef  srfi/srfi-18.scm
 ¢
	 £			 £		 £		 £	 		  g  nameg  uncaught-exception? C2RJ2f  h      ]4 5C        g  exc
		  g  filenamef  srfi/srfi-18.scm
 ¤
	 ¥		 ¥	/	 ¥		 ¥	 		  g  nameg  uncaught-exception-reason C3RM        h      ] C       g  obj
		  g  filenamef  srfi/srfi-18.scm
 ¦
	 §	 		  g  nameg  terminated-thread-exception? C1Rg!hijklkmno-W     hx      ] 4L>  "  G  4L>  "  G  4L>  "  G  4L>  "  G  4LL>  "  G  4L>  "  G  L 6{       g  filenamef  srfi/srfi-18.scm
 ±		 ²		 ³		' ´		9 µ		K ¶		_ ·		x ¸	 			x
   C]_RPon      hÐ   À  - 1 3 $  "  454545454>  "  G  4	 O 
54>  "  G  $  4>  "  G  "   4>  "  G  4>  "  G  C      ¸      g  thunk
		 Ê g  name		 Ê g  n		4 Ê g  lm		4 Ê g  lc		4 Ê g  sm		4 Ê g  sc		4 Ê g  t		f Ê  g  filenamef  srfi/srfi-18.scm
 º		 »		 »		 »	!	 ½		 ½		 ½		  ¾		$ ¾	)	& ¾		' ¿		+ ¿		- ¿		. À		2 À	)	4 À		4 »		? Â		Q Ã		f Ã		i Å	
	t Å	+	y Å	
  Æ	
  Æ	 ¢ Ç	
 ¶ È	
 		 Ê
   CR^PJ7p       h      ]4 56            g  thread
		  g  filenamef  srfi/srfi-18.scm
 Ë
	 Ì		 Ì	9	 Ì		 Ì	 		  g  nameg  thread-name CR^QJ7q   h      ]4 56            g  thread
		  g  filenamef  srfi/srfi-18.scm
 Î
	 Ð		 Ð	,	 Ð		 Ï	 		  g  nameg  thread-specific CR_QJ7r       h(   ·   ]44 5>  "  G  C     ¯       g  thread
		# g  obj		#  g  filenamef  srfi/srfi-18.scm
 Ò
	 Ó		 Ô		 Ô	-	 Ô		 Ó	 		#	  g  nameg  thread-specific-set! CR^RJ7stkmn 
       h   %  ]44 55$  ^4 >  "  G  4>  "  G  4>  "  G  4	>  "  G  "    C         g  thread
		} g  x		} g  smutex		"	v g  scond		"	v  g  filenamef  srfi/srfi-18.scm
 Ø
	 Ù		 Ú		 Ú	5	 Ú		 Ù		 Ù		 Û		 Û		" Ü		" Û		' Ý		; Þ		M ß		_ à	 		}  g  nameg  thread-start! CRu  h   f   ] 4>   "  G  C    ^       g  filenamef  srfi/srfi-18.scm
 ã
	 ã	 		
  g  nameg  thread-yield! CR*()vGHwIxyz{   h°     ]!445 54 5$  4 5"  %4 5$  	 "  4  54	4
554	4
è55
$  4>  "  G  "   
$  4>  "  G  "   C û      g  timeout
	 ¯ g  ct	 ¯ g  t		L ¯ g  secs		Z ¯ g  usecs		o ¯  g  filenamef  srfi/srfi-18.scm
 å
	 æ		 æ		 æ		 æ		 ç		 ç		 ç	&	# ç	#	( è		2 ç		7 è	%	< é		@ é	#	B é	3	D ê	#	I ë	#	J ì	#	L é		L æ		O í		R í		Z í		Z æ		] î		` î	 	g î	-	k î	*	m î	 	o î		o æ		u ï			y ï		z ï	  ð		  ð	  ð	 &	 ¯  g  nameg  thread-sleep! CR-,        h    z   ]445  >  "  G  L 6      r       g  obj
		  g  filenamef  srfi/srfi-18.scm
 û		 ü		 ü		 ü		 ý	 		   C   h   h   ] O L 6 `       g  continuation
		  g  filenamef  srfi/srfi-18.scm
 ú		 û	 		   C     h   h   ] O C      `       g  thunk
		
  g  filenamef  srfi/srfi-18.scm
 ù
 		
  g  nameg  wrap C|R}c~-WVM  h    V   ] 4L >  "  G  6    N       g  filenamef  srfi/srfi-18.scm
				!		! 		
   CVM  h   M   ] 6E       g  filenamef  srfi/srfi-18.scm
			* 		
   C      h`   Ñ   ]	4 545$  4 O >  "  G  "  4 >  "  G  4 >  "  G  C     É       g  thread
		[ g  current-handler			[  g  filenamef  srfi/srfi-18.scm

							
				
	4	
	H	 			[  g  nameg  thread-terminate! CR|Y.L   hh   ¼   ]4L L?4L 5 4L5$    $  "  4>  "  G  "   $  4>  "  G  "    C      ´       g  v
		b g  e		b  g  filenamef  srfi/srfi-18.scm
	
											#		)		.		I		J	 		b
   C h    ¶   - 1 3 H4 O 5KJB ®       g  thread
			 g  args			 g  thread-join-inner!		
	  g  filenamef  srfi/srfi-18.scm

				 			
  g  nameg  thread-join! CR_P       hP   ù   -  1  3  $   "  45$  4>  "  G  "   C      ñ       g  name
			J g  n	#	J g  m		#	J  g  filenamef  srfi/srfi-18.scm
#		$		$		$		%		%		&		'		!(		#%		#$		-)		.)	 			J


  g  nameg  
make-mutex CR^PJ8      h      ]4 56            g  mutex
		  g  filenamef  srfi/srfi-18.scm
+
	,		,	7	,		,	 		  g  nameg  
mutex-name CR^QJ8     h      ]4 56            g  mutex
		  g  filenamef  srfi/srfi-18.scm
.
	0		0	*	0		/	 		  g  nameg  mutex-specific CR_QJ8 h(   µ   ]44 5>  "  G  C     ­       g  mutex
		# g  obj		#  g  filenamef  srfi/srfi-18.scm
2
	3		4		4	+	4		3	 		#	  g  nameg  mutex-specific-set! CR        h8   ä   ]	4 5$  45$  CC4 5
$  CC   Ü       g  mutex
		5 g  owner			5  g  filenamef  srfi/srfi-18.scm
8
	9			9		:		;		;		;	#	#<		+<		/<		1<	&	4<	1 		5  g  nameg  mutex-state CR|k   h   M   ] LL @      E       g  filenamef  srfi/srfi-18.scm
B		
B	 		

   C.K   h   r   - 1 3 6 j       g  key
			 g  args			  g  filenamef  srfi/srfi-18.scm
C		C	( 			
   C   h   V   ] L LO 6     N       g  filenamef  srfi/srfi-18.scm
@	
	A		A	 		
   C       h    ³   - 1 3 H4 O 5KJB «       g  mutex
			 g  args			 g  mutex-lock-inner!		
	  g  filenamef  srfi/srfi-18.scm
>
	@		D	 			
  g  nameg  mutex-lock! CRn    h      - 1 3  @              g  mutex
			 g  args			  g  filenamef  srfi/srfi-18.scm
F
	G	 			
  g  nameg  mutex-unlock! C R_P      hH   ë   -  1  3  $   "  45 $  4>  "  G  "   C    ã       g  name
			D g  n		D g  m			D  g  filenamef  srfi/srfi-18.scm
M		N		N		N		O		O		O		N		'P		(P	 			D


  g  nameg  make-condition-variable C!R^PJ9    h   ¬   ]4 56     ¤       g  condition-variable
		  g  filenamef  srfi/srfi-18.scm
R
	S		U	*	S		S	 		  g  nameg  condition-variable-name C"R^QJ9   h   °   ]4 56     ¨       g  condition-variable
		  g  filenamef  srfi/srfi-18.scm
W
	X		Z	.	X		X	 		  g  nameg  condition-variable-specific C#R_QJ9       h(   Ï   ]44 5>  "  G  C     Ç       g  condition-variable
		# g  obj		#  g  filenamef  srfi/srfi-18.scm
\
	]		^		`		^		]	 		#	  g  nameg   condition-variable-specific-set! C$Rm        h      ]4 >  "  G  C  ~       g  cond
		  g  filenamef  srfi/srfi-18.scm
d
	e	 		  g  nameg  condition-variable-signal! C%R h      ]4 >  "  G  C         g  cond
		  g  filenamef  srfi/srfi-18.scm
h
	i	 		  g  nameg  condition-variable-broadcast! C&Ri(R       hH   ú   ]	 $  9 45$  
"  $   45$  
CCCC    ò       g  obj
		D g  co		% g  co	,	@  g  filenamef  srfi/srfi-18.scm
o
	p			p		q		q		q	"	q		q	0	)p		,r		,r		/r	"	9r		=r	0 		D  g  nameg  time? C)RJ) h(   ¸   ]4 5$         B@CC   °       g  time
		%  g  filenamef  srfi/srfi-18.scm
t
	u		u	"	u		u		v	
	v		!v		"v	 
		%  g  nameg  time->seconds C*RJvyx       h@   ú   ]	4 5$  +4 54544      B@55CC  ò       g  x
		> g  fx		<  g  filenamef  srfi/srfi-18.scm
x
	y		y	!	y		y		z		z		{		"|		%|		,|	,	6|	)	8|		:|		;{		 		>  g  nameg  seconds->time C+RC Á      g  m
		4  g  filenamef  srfi/srfi-18.scm		 
	5	[			;	[		=	[			A	[
	F	\		L	\		Q	\		Z	^
	_	^		g	^	&	l	^
x	`
z	f	(}	f	"	f
	g	%	g		g
	h	*	h	$	h
	i	!	i		i
	k	¢	k
£	l	¬	l
­	m	¶	m
·	n	"À	n
Å	r
k	s
l	v	u	v
	x
	}
 
§ 
	F 
	å  

 ¡
L ¢
 ¤
º ¦
 ¯
_ Ë
+ Î
* Ò
ø Ø
 ã
f å
+ ù
l

ü"
 Å+
!.
"2
#Ç8
&6>
&ðF
(>L
)R
*W
+\
+Ñd
,h
,n
-åo
.Ût
05x
 K	07
   C6 