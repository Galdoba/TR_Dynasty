package dynasty

/*
1. Roll Characteristics and determine Characteristic modifiers.
2.  a. Choose a Power Base.	+
    b. Gain Trait and Aptitude Modifiers.	+
3.  a. Choose a Dynasty Archetype.	+
	b. Determine Base Traits and Aptitudes.	+
	c. Gain First Generation bonuses.	+
	d. Determine Dynasty Boons and Hinders. 	+
4.  a. Determine Management Assets.
	b. Gain Management Asset Benefits.
5.  Calculate First Generation Values.
6. Move on to the Background and Historic Events process.

Мне нужна будет информация (возвращаясь к Столыпину):
- о Харреке, в том числе что известно о его взаимоотношениях с Олебом (и почему тот уверен, что сынок не хочет его подсидеть)
- о Рихарде (???), синдальском принце - максимально, в том числе, где он сейчас, что делает, чем вообще живёт, какие у него больные места
- о расстановке сил на основных торговых путях, может какие-то имена, их сильные/слабые стороны, узкие места (и желательно сами пути ещё раз указать, я уже плохо помню, что там где проходит((()
Меня не оставляет подкинутая идея о кораблях внутри корабля 😑 Можно два истребителя назвать Клык и Коготь и нанять на них отряд боевых котиков и варгров 😁 а прыжковый назвать База, как в каких-нибудь шпионских фильмах :D

Харрик:
Первый, старший и единственный сын Олеба. Признаный наследником. Прирожденный лидер способный принять быстрое ввешенное решение, склонный к поисков компромиссов и не принимающий бессмысленного насилия.
В 1086 году, когда Дринакс из-за неполадок с гидропоникой потерял почти весь урожай, поприказу своего отца возглавил рейд на Асим. Целью которого было решения продовольственного вопроса.
Однако во время высадки его шатл был сбит и Харрек получил тяжелые ранения. Впав в бешенство Олем отдал приказ уничтожить руководство Асима и отправить едва живого принца в Башню (научный центр Дринакса), потребовав вылечить сына любыми средствами.
Несколько месяцев назад Харек вышел из комы (хотя никто уже не может сказать чего в нем больше - киберимплантов или клонированых тканей). Несколько месяцев Харрик пытался осознать новый мир в котором он оказался. Его новое тело по слухам обладает нечеловеческими способностями, он больше не наследник трона (во избежании волнений, Олеб назначил наследницей Рао, которая имеет как возможности так и амбиции к управлению).
Сейчас вокруг Харрика собирается все больше знати Дринакса, считающей что Дринакс должен сконцентрироваться на том чтобы восстановить экологию планеты, а не ввязываться в игры внешней политики. Сам Харрик не согласен с планом Олеба - взять под контроль торговые пути между Империей и Иерархатом используя пиратство. Однако к тому моменту когда, он очнулся план уже пришел в движение и Харрек не пошел против воли отца.

Рихард:
О нем известно не многое. На политической сцене он появился около полутора лет назад, однако очень быстро начал получать поддержку в торговых кругах и в широких массах на нескольких планетах.
Его основные призывы сводятся к тому, что миры Пограничья и Пылевого Пояса должны объединиться для обеспечения безопасности как населения планет так и торговых путей между Империй, Иерархатом и Флорианской Лигой. На фоне всё учащающихся рейдов Ихотей, Рихард плучает все больше поддержки среди народных масс.
У Рихарда нет основной резиденции, потому что он часто перемещается между мирами Пограничья и Пылевого Пояса.

Общий расклад:

Торговля:
2 торговых путя связываю Империум и Иерархат. Путь-2 и Путь-3 получили свое название от рейтинга прыжковых двигателей необходимых для их преодоления.
Путь-3:
Fist -> Wildeman -> Acrid -> Techworld -> Paal -> Tyokh
Путь-2:
Fist -> Wildeman -> Cordan -> Argona -> Sperle -> Techworld -> Paal -> Tyokh
Оба государства заинтересованы в торговле, однако оба стараются не провоцировать друг друга отправляя флот на обоспечение безопасности торговли. В свою очередь пираты стараются не провоцировать обе силы и грабят в основном мелких и независимых торговцев. Такая ситуация устраивает всех... кроме независимых торговцев и местного населения.

Пираты:
Около 30 лет назад из Империума бежал бывший флотский офицер Кайл Дарокин. Он осел на планете Тэв (Theev) и сколотил свою банду. Сейчас Адмирал Дарокин возглавляет совет баронов на Тэве и является самым могущественным пиратским лидером во всем секторе.
Тэв держит в тонусе торговлю между Империумом и Иерархатом, и практически контролирует торговлю межде Империумом и Флорианской Лигой.

Независимые миры:
Большая часть миров предоставленна сама себе и живет лишь за счет того что обеспечивает торговый трафик. Однако большая часть миров не способна сколько-нибудь эффективно контролировать свои системы за пределами радиуса действий планетарных и орбитальных орудий.
Этим активно пользуются Пираты, Контрабандисты и Ихотей.

Ихотей:
"Младшие сыновья" Асланских кланов. Формируются в банды в поисках богатства и славы. Договор о границах между Империумом и Иерархатом не позволяет кланам захватывать новые земли (однако срок его действия истекает в 1107 году). Однако официально Ихотей не являются представителями Иерархата и не захватывают новые миры. Последние годы распространилась практика основывания асланских поселений на формально независимых мирах. И Империуму приходится с этим положением дел.
Ситуация с Ихотей накаляется с увеличением прямых стычек между людьми и асланами в независимых мирах. Так же отмечается особая дикость и жестокость с которой асланы вырезали несколько поселений на Тире (Tyr). Правительство Тира открыто обвиняет Акис (Acis) в пособничестве Ихотей, Акис в свою очередь настаивает на том что рейду Ихотей это общая проблема сектора и в одиночку они не могут защищать весь сектор от банд которые проникают в Пылевой Пояс через их систему. Ходят слухи что Тир уже в процессе мобилизации своего флота.
Кроме этого отмечается увеличение Ихотейских лагерей в системах Кхусай (Khusai) и Ахвокиал (Akhwohkyal). Однако отсуствие массовых рейдов и захватов территорий связывают с неспособностью банд выстроить четкую иерархию. Но все может резко измениться если появится Лидер который сможет направить эту лавину.

ТУДУ:
выдать информацию по
Хареку (эффект 4)
Рихарду (Эффект 0)
Расстановке сил в секторе (эффект 3)
*/
