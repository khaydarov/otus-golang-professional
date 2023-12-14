package hw03frequencyanalysis

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var testCases = []struct {
	text            string
	expected        []string
	testDescription string
}{
	{
		"",
		[]string{},
		"no words in empty string",
	},
	{
		`Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`,
		[]string{
			"а",         // 8
			"он",        // 8
			"и",         // 6
			"ты",        // 5
			"что",       // 5
			"в",         // 4
			"его",       // 4
			"если",      // 4
			"кристофер", // 4
			"не",        // 4
		},
		"adventures of Winnie the Pooh in cyrillic",
	},
	{
		`English is the most widely spoken language in the world, with over 1.5 billion speakers. 
			It is the official language of 53 countries and is used in many other countries as a lingua franca.
			English is also the language of science, technology, and business. As a result, it is essential for anyone 
			who wants to succeed in the globalized world to be able to speak and understand English.`,
		[]string{
			"is",        // 5
			"the",       // 5
			"and",       // 3
			"english",   // 3
			"in",        // 3
			"language",  // 3
			"to",        // 3
			"a",         // 2
			"as",        // 2
			"countries", // 2
		},
		"random english text",
	},
	{
		"English is the most widely spoken language in the world",
		[]string{
			"the",      // 2
			"english",  // 1
			"in",       // 1
			"is",       // 1
			"language", // 1
			"most",     // 1
			"spoken",   // 1
			"widely",   // 1
			"world",    // 1
		},
		"short random text where distinct words are less than 10",
	},
	{
		"Hello World!",
		[]string{
			"hello",
			"world",
		},
		"simple hello world test",
	},
}

func TestTop10(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.testDescription, func(t *testing.T) {
			require.Equal(t, testCase.expected, Top10(testCase.text))
		})
	}
}
