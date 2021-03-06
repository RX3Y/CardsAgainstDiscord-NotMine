package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "40-blanks",
		Description: "Deck of 40 blank response cards",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `%s may pass, but %s will last forever.`},
			&PromptCard{Prompt: `%s will never be the same after %s.`},
			&PromptCard{Prompt: `"This is madness." "No, THIS IS %s!"`},
			&PromptCard{Prompt: `2 AM in the city that never sleeps. The door swings open and she walks in, legs up to here. Something in her eyes tells me she's looking for %s.`},
			&PromptCard{Prompt: `Adventure. Romance. %s. From Paramount Pictures, "%s."`},
			&PromptCard{Prompt: `And today's soup is Cream of %s.`},
			&PromptCard{Prompt: `And would you like those buffalo wings mild, hot, or %s?`},
			&PromptCard{Prompt: `Armani suit: $1,000. Dinner for two at that swanky restaurant: $300. The look on her face when you surprise her with %s: priceless.`},
			&PromptCard{Prompt: `As king, how will I keep the peasants in line? %s`},
			&PromptCard{Prompt: `Behind every powerful man is %s.`},
			&PromptCard{Prompt: `Come to Dubai, where you can relax in our world-famous spas, experience the nightlife, or simply enjoy %s by the poolside.`},
			&PromptCard{Prompt: `Dammit Gary. You can't just solve every problem with %s.`},
			&PromptCard{Prompt: `Dear Leader Kim Jong-un, our village praises your infinite wisdom with a humble offering of %s.`},
			&PromptCard{Prompt: `Do not fuck with me! I am literally %s right now.`},
			&PromptCard{Prompt: `Do the Dew with our most extreme flavor yet! Get ready for Mountain Dew %s!`},
			&PromptCard{Prompt: `Do you lack energy? Does it sometimes feel like the whole world is %s ? Ask your doctor about Zoloft.®`},
			&PromptCard{Prompt: `Don't forget! Beginning this week, Casual Friday will officially become "%s Friday."`},
			&PromptCard{Prompt: `Don't worry kid. It gets better. I've been living with %s for 20 years.`},
			&PromptCard{Prompt: `Every step towards %s gets me a little bit closer to %s.`},
			&PromptCard{Prompt: `Everybody join hands and close your eyes. Do you sense that? That's the presence of %s in this room.`},
			&PromptCard{Prompt: `Forget everything you know about %s, because now we've supercharged it with %s!`},
			&PromptCard{Prompt: `Get ready for the movie of the summer! One cop plays by the book. The other's only interested in one thing: %s.`},
			&PromptCard{Prompt: `Having the worst day EVER. #%s`},
			&PromptCard{Prompt: `Heed my voice, mortals! I am the god of %s , and I will not tolerate %s!`},
			&PromptCard{Prompt: `Help me doctor, I've got %s in my butt!`},
			&PromptCard{Prompt: `Here at the Academy for Gifted Children, we allow students to explore %s at their own pace.`},
			&PromptCard{Prompt: `Hi MTV! My name is Kendra, I live in Malibu, I'm into %s, and I love to have a good time.`},
			&PromptCard{Prompt: `Hi, this is Jim from accounting. We noticed a $1,200 charge labeled "%s." Can you explain?`},
			&PromptCard{Prompt: `Honey, I have a new role-play I want to try tonight! You can be %s, and I'll be %s.`},
			&PromptCard{Prompt: `How am I compensating for my tiny penis?%s`},
			&PromptCard{Prompt: `I am become %s, destroyer of %s!`},
			&PromptCard{Prompt: `I don't mean to brag, but they call me the Micheal Jordan of %s.`},
			&PromptCard{Prompt: `I have a strict policy. First date, dinner. Second date, kiss. Third date, %s.`},
			&PromptCard{Prompt: `I work my ass off all day for this family, and this is what I come home to? %s!?`},
			&PromptCard{Prompt: `I'm Miss Tennessee, and if I could make the world better by changing one thing, I would get rid of %s.`},
			&PromptCard{Prompt: `I'm pretty sure I'm high right now, because I"m absolutely mesmerized by %s.`},
			&PromptCard{Prompt: `I'm sorry, Mrs. Chen, but there was nothing we could do. At 4:15 this morning, your son succumbed to %s.`},
			&PromptCard{Prompt: `I'm sorry, sir, but we don't allow %s at the country club.`},
			&PromptCard{Prompt: `If you can't handle %s, you'd better stay away from %s.`},
			&PromptCard{Prompt: `If you had to describe the Card Czar, using only one of the cards in your hand, which one would it be? %s`},
			&PromptCard{Prompt: `In his farewell address, George Washington famously warned Americans about the dangers of %s.`},
			&PromptCard{Prompt: `In his new action comedy, Jackie Chan must fend off ninjas while also dealing with %s.`},
			&PromptCard{Prompt: `In return for my soul, the Devil promised me %s, but all I got was %s.`},
			&PromptCard{Prompt: `In the beginning, there was %s. And the Lord said, "Let there be %s."`},
			&PromptCard{Prompt: `It lurks in the night. It hungers for flesh. This summer, no one is safe from %s.`},
			&PromptCard{Prompt: `James is a lonely boy. But when he discovers a secret door in his attic, he meets a magical new friend: %s.`},
			&PromptCard{Prompt: `Life's pretty tough in the fast lane. That's why I never leave the house without %s.`},
			&PromptCard{Prompt: `Listen Gary, I like you. But if you want that corner office, you're going to have to show me %s.`},
			&PromptCard{Prompt: `Man, this is bullshit. Fuck %s.`},
			&PromptCard{Prompt: `My grandfather worked his way up from nothing. When he came to this country, all he had was the shoes on his feet and %s.`},
			&PromptCard{Prompt: `Now in bookstores: "The Audacity of %s" by Barack Obama.`},
			&PromptCard{Prompt: `Oprah's book of the month is "%s For %s: A Story of Hope."`},
			&PromptCard{Prompt: `Patient presents with %s. Likely a result of %s.`},
			&PromptCard{Prompt: `Puberty is a time of change. You might notice hair growing in new places. You might develop an intrest in %s. This is normal.`},
			&PromptCard{Prompt: `She's up all night for good fun. I'm up all night for %s.`},
			&PromptCard{Prompt: `The Japanese have developed a smaller, more efficient version of %s.`},
			&PromptCard{Prompt: `The six things I could never do without: oxygen, Facebook, chocolate, Netflix, friends, and %s LOL!`},
			&PromptCard{Prompt: `This is America. If you don't work hard, you don't succeed. I don't care if you're black, white, purple, or %s.`},
			&PromptCard{Prompt: `This is the prime of my life. I'm young, hot, and full of %s.`},
			&PromptCard{Prompt: `This year's hottest album is "%s" by %s.`},
			&PromptCard{Prompt: `To become a true Yanomami warrior, you must prove that you can withstand %s without crying out.`},
			&PromptCard{Prompt: `Tonight we will have sex. And afterwards, If you'd like, a little bit of %s.`},
			&PromptCard{Prompt: `We never did find %s, but along the way we sure learned a lot about %s.`},
			&PromptCard{Prompt: `Well if %s is good enough for %s, it's good enough for me.`},
			&PromptCard{Prompt: `Well what do you have to say for yourself, Casey? This is the third time you've been sent to the principal's office for %s.`},
			&PromptCard{Prompt: `Wes Anderson's new film tells the story of a precocious child coming to terms with %s.`},
			&PromptCard{Prompt: `What killed my boner? %s`},
			&PromptCard{Prompt: `What's fun until it gets weird? %s`},
			&PromptCard{Prompt: `What's making things awkward in the sauna? %s and %s`},
			&PromptCard{Prompt: `When I was a kid, we used to play Cowboys and %s.`},
			&PromptCard{Prompt: `WHOOO! God damn I love %s!`},
			&PromptCard{Prompt: `Why am I broke? %s`},
			&PromptCard{Prompt: `Why won't you make love to me anymore? Is it %s?`},
			&PromptCard{Prompt: `Y'all ready to get this thing started? I'm Nick Cannon, and this is America's Got %s.`},
			&PromptCard{Prompt: `Yo' mama's so fat she %s!`},
			&PromptCard{Prompt: `You are not alone. Millions of Americans struggle with %s every day.`},
			&PromptCard{Prompt: `You know, once you get past %s, %s ain't so bad.`},
			&PromptCard{Prompt: `You Won't Believe These 15 Hilarious %s Bloopers!`},
			&PromptCard{Prompt: `You've seen the bearded lady! You've seen the ring of fire! Now, ladies and gentlemen, feast your eyes upon %s!`},
		},
		Responses: []ResponseCard{
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
			`%blank`,
		},
	}

	AddPack(pack)
}
