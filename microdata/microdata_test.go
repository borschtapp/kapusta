package microdata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonLd(t *testing.T) {
	recipeUrl := "https://allrecipes.com/recipe/231495/texas-boiled-beer-shrimp/"

	data, err := ParseURL(recipeUrl)
	assert.Nil(t, err)

	b, err := json.MarshalIndent(data, "", "  ")
	assert.Equal(t, `{
  "items": [
    {
      "type": [
        "BreadcrumbList"
      ],
      "properties": {
        "@context": [
          "http://schema.org"
        ],
        "itemListElement": [
          {
            "type": [
              "ListItem"
            ],
            "properties": {
              "item": [
                {
                  "type": [],
                  "properties": {
                    "@id": [
                      "https://www.allrecipes.com/"
                    ],
                    "name": [
                      "Home"
                    ]
                  }
                }
              ],
              "position": [
                1
              ]
            }
          },
          {
            "type": [
              "ListItem"
            ],
            "properties": {
              "item": [
                {
                  "type": [],
                  "properties": {
                    "@id": [
                      "https://www.allrecipes.com/recipes/"
                    ],
                    "name": [
                      "Recipes"
                    ]
                  }
                }
              ],
              "position": [
                2
              ]
            }
          },
          {
            "type": [
              "ListItem"
            ],
            "properties": {
              "item": [
                {
                  "type": [],
                  "properties": {
                    "@id": [
                      "https://www.allrecipes.com/recipes/93/seafood/"
                    ],
                    "name": [
                      "Seafood"
                    ]
                  }
                }
              ],
              "position": [
                3
              ]
            }
          },
          {
            "type": [
              "ListItem"
            ],
            "properties": {
              "item": [
                {
                  "type": [],
                  "properties": {
                    "@id": [
                      "https://www.allrecipes.com/recipes/412/seafood/shellfish/"
                    ],
                    "name": [
                      "Shellfish"
                    ]
                  }
                }
              ],
              "position": [
                4
              ]
            }
          },
          {
            "type": [
              "ListItem"
            ],
            "properties": {
              "item": [
                {
                  "type": [],
                  "properties": {
                    "@id": [
                      "https://www.allrecipes.com/recipes/430/seafood/shellfish/shrimp/"
                    ],
                    "name": [
                      "Shrimp"
                    ]
                  }
                }
              ],
              "position": [
                5
              ]
            }
          }
        ]
      }
    },
    {
      "type": [
        "Recipe"
      ],
      "properties": {
        "@context": [
          "http://schema.org"
        ],
        "aggregateRating": [
          {
            "type": [
              "AggregateRating"
            ],
            "properties": {
              "bestRating": [
                "5"
              ],
              "itemReviewed": [
                "Texas Boiled Beer Shrimp"
              ],
              "ratingCount": [
                24
              ],
              "ratingValue": [
                3.75
              ],
              "reviewCount": [
                21
              ],
              "worstRating": [
                "1"
              ]
            }
          }
        ],
        "author": [
          {
            "type": [
              "Person"
            ],
            "properties": {
              "name": [
                "waldapfel79"
              ]
            }
          }
        ],
        "cookTime": [
          "P0DT0H5M"
        ],
        "datePublished": [
          "2013-02-22T22:24:10.000Z"
        ],
        "description": [
          "Family favorite. Serve with lemon wedges."
        ],
        "image": [
          {
            "type": [
              "ImageObject"
            ],
            "properties": {
              "height": [
                3456
              ],
              "url": [
                "https://imagesvc.meredithcorp.io/v3/mm/image?url=https%3A%2F%2Fimages.media-allrecipes.com%2Fuserphotos%2F1068356.jpg"
              ],
              "width": [
                4608
              ]
            }
          }
        ],
        "mainEntityOfPage": [
          "https://www.allrecipes.com/recipe/231495/texas-boiled-beer-shrimp/"
        ],
        "name": [
          "Texas Boiled Beer Shrimp"
        ],
        "nutrition": [
          {
            "type": [
              "NutritionInformation"
            ],
            "properties": {
              "calories": [
                "168 calories"
              ],
              "carbohydrateContent": [
                "4.2 g"
              ],
              "cholesterolContent": [
                "230.4 mg"
              ],
              "fatContent": [
                "1.3 g"
              ],
              "proteinContent": [
                "25.3 g"
              ],
              "saturatedFatContent": [
                "0.3 g"
              ],
              "sodiumContent": [
                "269.4 mg"
              ]
            }
          }
        ],
        "prepTime": [
          "P0DT0H5M"
        ],
        "recipeCategory": [
          "Shrimp Recipes"
        ],
        "recipeIngredient": [
          "2 (12 fluid ounce) cans or bottles light beer",
          "2 tablespoons dry crab boil",
          "2 pounds large shrimp"
        ],
        "recipeInstructions": [
          {
            "type": [
              "HowToStep"
            ],
            "properties": {
              "text": [
                "Bring beer with dry crab boil to a boil in a large pot. Add shrimp to boiling beer and place a cover on the pot. Bring the beer again to a boil, reduce heat to medium-low, and cook at a simmer for 5 minutes. Remove pot from heat and leave shrimp steeping in the beer another 2 to 3 minutes; drain. Serve immediately.\n"
              ]
            }
          }
        ],
        "recipeYield": [
          "2 pounds shrimp"
        ],
        "review": [
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "Dave Mcvety"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2018-01-19T23:34:00.720Z"
              ],
              "reviewBody": [
                "did not make it , did not have to. if you want little round rubber balls cook it this way. a waste of good shrimp. shrimp should NEVER be boiled. if you want plump tender shrimp, try this method. it never fails no matter the size or amount. in a large stock pot add 2 quarts of liquid, i use water, for every pound of shrimp. add a tbls. of salt for every 2 quarts of water.add 4or5 bay leaves. bring to a rapid boil and take it off the heat. put lid on pot and wait 2 mins. the water should be about 190 degree's now. dump in the shrimp, stir a bit, put lid on pot. exactly 3 mins later drain in colinder and immediatly plunge into ice water stirring until the shrimp are cold. spin in a salad drier or pat dry with paper towels. i like my shrimp cocktail shrimp really cold so i cover and put in the fridge but they can be served right away. try this method and you won't cook them any other way again.   David."
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      1
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "AcaCandy"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2013-11-17T19:58:32.057"
              ],
              "reviewBody": [
                "I didn't have the crab boil, so I had to substitute spices I had on hand. I used a blend of minced onion/garlic and some dried parsley and dill. I also didn't want to overcook, so as soon as the shrimp started to boil and turn pink (I used shell on shrimp) I pulled them off the heat, and drained immediately. I then put them in the freezer to cool down to serve as a cold appetizer with some home made cocktail sauce (organic ketchup and horseradish mix)."
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      5
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "BigShotsMom"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2013-03-01T18:27:24.93"
              ],
              "reviewBody": [
                "This was well received at my house.  The only thing I did differently was to remove the pot from the heat as soon as it came to a boil after adding the shrimp.  I didn't want to over cook the shrimp.  I left them in the hot beer for about 5 minutes.  Thanks for sharing!"
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      4
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "roughboy50"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2014-12-14T19:15:58.460Z"
              ],
              "reviewBody": [
                "Very similar to what i do,except I add a teaspoon of vinegar to it so they peel easier,I also add a tablespoon of caraway seeds to the boil."
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      4
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "Ardeen Rentsch"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2018-01-21T10:03:06.750Z"
              ],
              "reviewBody": [
                "Honestly, Florida native with LOTS of shrimp experience.  No matter WHAT you cook them in, bring liquid to a boil, turn off heat, and let shrimp sit in hot liquid about 5 or so minutes.  Otherwise you get rubber...."
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      2
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "Sue Wright"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2018-01-20T03:00:39.807Z"
              ],
              "reviewBody": [
                "We call this peel 'em and eat 'em. I use Old Bay, peppercorns and bay leaf. Bring beer and spices to boil, throw shrimp in, bring to boil, turn off. Let sit for 5 minutes. I prefer them warm with homemade cocktail sauce (ketchup, lemon, horseradish). For fun, cover the table with newspaper and thrown the peels on that! (I also serve green salad and good La Brea bread.)"
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      4
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "good eats"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2018-01-19T21:28:40.910Z"
              ],
              "reviewBody": [
                "Ridiculously overcooked."
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      2
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "Gail"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2018-01-20T03:49:30.110Z"
              ],
              "reviewBody": [
                "Bring the liquid to a boil,  add the shrimp, and as soon as it starts to boil again, take it off the heat and drain the liquid off.  \nNEVER let the shrimp stay in the boiling liquid any longer.  \nThis way you will have the true texture and flavour of the shrimp.  Bon appetite !"
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      3
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "Alisha"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2015-05-08T17:20:17.390Z"
              ],
              "reviewBody": [
                "Simple, easy and delicious.  At the recommendation of another reviewer, we also added 1 teaspoon of vinegar. The shrimp peeled very easily and tasted yummy!  Will definitely use this recipe again."
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      5
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "Jim Madison"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2018-01-26T01:26:45.100Z"
              ],
              "reviewBody": [
                "It was definitely right up there with the Finger Lickin  good"
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      5
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "Loulala"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2016-01-12T01:07:30.027Z"
              ],
              "reviewBody": [
                "A favorite dish of mine since college days, generations ago. I make mine with a small scoop of pickling spices as well. We call them \"peel and eat\" shrimp at my table, served up with some spicy cocktail sauce enhanced with added horseradish."
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      5
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "Carole Garretson Peddy"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2018-01-24T22:55:23.400Z"
              ],
              "reviewBody": [
                "Great base recipe!  I turned it into a shrimp boil by adding red potatoes , corn on the cob and beef smoked sausage so it made a delicious soup. Just be sure to taste it after adding everything else to see if you need to add more seasoning. Oh, and I used chicken broth in place of water like I do with most all things. We ate every single drop of it!!  Will definitely make again and again!"
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      5
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "lliampitt"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2018-01-23T03:00:01.29"
              ],
              "reviewBody": [
                "This is a pretty good shrimp recipe although it turns out kind of tough."
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      4
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "Teddy Lassiter"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2018-09-03T18:21:12.757Z"
              ],
              "reviewBody": [
                "Was delicious"
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      5
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "Sally558"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2018-01-21T00:22:12.830Z"
              ],
              "reviewBody": [
                "I love shrimp.  Made a half order, only 2 of us.  Great taste but followed directions and way over cooked.  Well make again but less cooking time."
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      3
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "RobynAnne"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2018-03-18T01:07:12.34"
              ],
              "reviewBody": [
                "Easy and DELICIOUS! When I lived in Seattle and purchased fresh sea food downtown at the market, this is exactly the way they recommended fresh shrimp be prepared."
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      5
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "Joan Landis Peters"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2018-02-10T00:31:14.613Z"
              ],
              "reviewBody": [
                "I will never make it again, No favor"
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      1
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "Deborah Dibble"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2018-01-27T00:42:19.903Z"
              ],
              "reviewBody": [
                "I used Old Bay seasoning because that's what I had on hand. It's the best shrimp I ever made. Tastes good hot or cold."
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      5
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "Roy"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2018-01-24T23:10:17.420Z"
              ],
              "reviewBody": [
                "Was good. Will make again."
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      4
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "Teresa Rosser Welch"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2018-02-14T20:39:22.340Z"
              ],
              "reviewBody": [
                "Simple ,tasty recipe."
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      5
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          },
          {
            "type": [
              "Review"
            ],
            "properties": {
              "author": [
                {
                  "type": [
                    "Person"
                  ],
                  "properties": {
                    "name": [
                      "truegravity43"
                    ]
                  }
                }
              ],
              "datePublished": [
                "2018-01-21T14:56:50.247Z"
              ],
              "reviewBody": [
                "You didn’t specify that you need to get local shrimp. Texas has local shrimp. Most shrimp you buy frozen from the grocery store is from Malaysia."
              ],
              "reviewRating": [
                {
                  "type": [
                    "Rating"
                  ],
                  "properties": {
                    "bestRating": [
                      "5"
                    ],
                    "ratingValue": [
                      1
                    ],
                    "worstRating": [
                      "1"
                    ]
                  }
                }
              ]
            }
          }
        ],
        "totalTime": [
          "P0DT0H10M"
        ]
      }
    }
  ]
}`, string(b))
}

func TestMicrodata(t *testing.T) {
	recipeUrl := "https://www.kuchnia-domowa.pl/przepisy/dania-glowne/595-krokiety-z-pieczarkami-jajkiem-zoltym-serem"

	data, err := ParseURL(recipeUrl)
	assert.Nil(t, err)

	b, err := json.MarshalIndent(data, "", "  ")
	assert.Equal(t, `{
  "items": [
    {
      "type": [
        "https://schema.org/BreadcrumbList"
      ],
      "properties": {
        "itemListElement": [
          {
            "type": [
              "https://schema.org/ListItem"
            ],
            "properties": {
              "item": [
                {
                  "type": [
                    "https://schema.org/WebPage"
                  ],
                  "properties": {
                    "name": [
                      "Przepisy"
                    ],
                    "url": [
                      "/przepisy"
                    ]
                  },
                  "id": "https://www.kuchnia-domowa.pl/przepisy"
                }
              ],
              "position": [
                "1"
              ]
            }
          },
          {
            "type": [
              "https://schema.org/ListItem"
            ],
            "properties": {
              "item": [
                {
                  "type": [
                    "https://schema.org/WebPage"
                  ],
                  "properties": {
                    "name": [
                      "Dania główne"
                    ],
                    "url": [
                      "/przepisy/dania-glowne"
                    ]
                  },
                  "id": "https://www.kuchnia-domowa.pl/przepisy/dania-glowne"
                }
              ],
              "position": [
                "2"
              ]
            }
          },
          {
            "type": [
              "https://schema.org/ListItem"
            ],
            "properties": {
              "item": [
                {
                  "type": [
                    "https://schema.org/WebPage"
                  ],
                  "properties": {
                    "name": [
                      "Krokiety z pieczarkami, jajkiem i żółtym serem"
                    ]
                  }
                }
              ],
              "position": [
                "3"
              ]
            }
          }
        ]
      }
    },
    {
      "type": [
        "http://schema.org/Recipe"
      ],
      "properties": {
        "aggregateRating": [
          {
            "type": [
              "https://schema.org/AggregateRating"
            ],
            "properties": {
              "bestRating": [
                "5"
              ],
              "ratingCount": [
                "3"
              ],
              "ratingValue": [
                "4"
              ],
              "worstRating": [
                "1"
              ]
            }
          }
        ],
        "articleBody": [
          "\n<img id=\"article-img-1\" class=\"article-img article-img-1\" title=\"Krokiety z pieczarkami, jajkiem i żółtym serem\" itemprop=\"image\" src=\"//static.kuchnia-domowa.pl/images/content/595/krokiety_z_pieczarkami_jajkiem_serem.jpg\" alt=\"Krokiety z pieczarkami, jajkiem i żółtym serem\" />\n\nKrokiety z farszem z usmażonych pieczarek, jajek na twardo, natki pietruszki i startego, żółtego sera. Bardzo smaczny obiad, prosty do przygotowania, ale nieco pracochłonny. Krokiety można podawać z filiżanką czerwonego barszczu.\n\n\n\t\n\n\n\n\n\n\nSkładniki na ok. 8- 9 sztuk:\n\nCiasto naleśnikowe:\n\n250 ml mleka\n250 ml gazowanej wody mineralnej*\n2 jajka\n225 g mąki pszennej\n½ łyżeczki soli\n\nFarsz:\n\n400 g pieczarek\n1 cebula\n2 łyżki oleju np. rzepakowego\n4 jajka\n100- 150 g startego sera żółtego\n2 łyżki posiekanej natki pietruszki\nsól, pieprz\n\nPanierka:\n\n2 jajka\nbułka tarta\nolej do smażenia\n\n\n\n\t\n\n.adsense_content2_rectangle2{display:inline-block}@media (max-width:740px){.adsense_content2_rectangle2{display:none !important}}\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nSposób przygotowania:\n\nPrzygotować farsz: Cebulę obrać i pokroić w kostkę. Pieczarki oczyścić i pokroić w plasterki lub posiekać drobno.\nNa 2 łyżkach oleju podsmażyć cebulę, aż się zeszkli. Dodać pieczarki i smażyć do momentu wyparowania całej wody i lekkiego przyrumienienia się pieczarek. Odstawić na bok do wystygnięcia.\nW międzyczasie ugotować jajka na twardo, zalać zimną wodą i po ostudzeniu obrać. Następnie posiekać i dodać do chłodnych pieczarek. Dodać posiekaną natkę pietruszki. Dobrze przyprawić solą i pieprzem. (Farsz powinien być wyrazisty w smaku). Wmieszać ser żółty.\nPrzygotować ciasto naleśnikowe: Mleko, wodę i jajka wymieszać np. trzepaczką. Dodać mąkę i zamieszać tak, aby nie powstały grudki. (Można zmiksować blenderem). Dodać sól i wymieszać.\nNa patelni o średnicy 28 cm usmażyć naleśniki na odrobinie oleju. Gotowe układać na talerzu, jeden na drugim.\nFarsz rozłożyć na naleśnikach, zostawiając brzegi wolne. Założyć boki naleśnika na farsz i zwinąć w rulon.\n\n<img class=\"article-img article-img-1\" src=\"//static.kuchnia-domowa.pl/images/content/595/2.jpg\" alt=\"2\" />\n<img class=\"article-img article-img-1\" src=\"//static.kuchnia-domowa.pl/images/content/595/3.jpg\" alt=\"3\" />\n<img class=\"article-img article-img-1\" src=\"//static.kuchnia-domowa.pl/images/content/595/4.jpg\" alt=\"4\" />\n\nJajka rozkłócić w głębokim talerzu. Do drugiego talerza wsypać bułkę tartą.\nNaleśniki obtaczać w jajku i następnie w bułce tartej.\nSmażyć na rozgrzanym na patelni oleju z każdej strony, aż do zarumienienia się bułki.\n\n\nUwaga:\n* W przypadku braku wody gazowanej można użyć zwykłej.\n\n\n\n\t\n\n\n\n\n\n\n\n\n\n\n\n\n\nCałkowity czas przygotowania:\n\n\nok. 2 godzin\n\n\n\n\n\nLiczba porcji:\n\n\nok. 4 (po 2 krokiety)\n\n\n\n\nTrudność:\nłatwa\n\n\n\nKoszt:\nśredni\n\n\n\n\n<img id=\"article-img-2\" class=\"article-img article-img-2\" title=\"Krokiety z pieczarkami, jajkiem i żółtym serem\" itemprop=\"image\" src=\"//static.kuchnia-domowa.pl/images/content/595/krokiety_z_pieczarkami_jajkiem_serem_2.jpg\" alt=\"Krokiety z pieczarkami, jajkiem i żółtym serem\" /> <img id=\"article-img-3\" class=\"article-img article-img-3\" title=\"Krokiety z pieczarkami, jajkiem i żółtym serem\" itemprop=\"image\" src=\"//static.kuchnia-domowa.pl/images/content/595/krokiety_z_pieczarkami_jajkiem_serem_3.jpg\" alt=\"Krokiety z pieczarkami, jajkiem i żółtym serem\" />Powiązane przepisy\n\n\n\n\n<img src=\"//static.kuchnia-domowa.pl/images/resized/images/content/88/krokiety_z_pieczarkami_60_60.jpg\"  alt=\"Krokiety z pieczarkami\" title=\"Krokiety z pieczarkami\" width=\"60\" height=\"60\" class=\"img-thumbnail img-polaroid\" />\nKrokiety z pieczarkami\n\n\nTags: Na ciepłoGrzybyMąkaJajkoMlekoSmak neutralnyMąka pszennaObiadKolacjaBoże NarodzenieUrodziny , RocznicaImpreza , prywatkaNa ślub , weselePolskaWok, patelniaBez mięsaWegetariańskieŁatwe\n"
        ],
        "author": [
          {
            "type": [
              "https://schema.org/Organization"
            ],
            "properties": {
              "name": [
                "Kuchnia Domowa"
              ]
            }
          }
        ],
        "creator": [
          {
            "type": [
              "https://schema.org/Organization"
            ],
            "properties": {
              "name": [
                "Kuchnia Domowa"
              ]
            }
          }
        ],
        "description": [
          "\nKrokiety z farszem z usmażonych pieczarek, jajek na twardo, natki pietruszki i startego, żółtego sera. Bardzo smaczny obiad, prosty do przygotowania, ale nieco pracochłonny. Krokiety można podawać z filiżanką czerwonego barszczu.\n"
        ],
        "image": [
          "https://static.kuchnia-domowa.pl/images/content/595/krokiety_z_pieczarkami_jajkiem_serem.jpg",
          "https://static.kuchnia-domowa.pl/images/content/595/krokiety_z_pieczarkami_jajkiem_serem_2.jpg",
          "https://static.kuchnia-domowa.pl/images/content/595/krokiety_z_pieczarkami_jajkiem_serem_3.jpg"
        ],
        "inLanguage": [
          "pl-PL"
        ],
        "mainEntityOfPage": [
          "https://www.kuchnia-domowa.pl/przepisy/dania-glowne/595-krokiety-z-pieczarkami-jajkiem-zoltym-serem/"
        ],
        "name": [
          "Krokiety z pieczarkami, jajkiem i żółtym serem",
          "\n\t\t\t\tKrokiety z pieczarkami, jajkiem i żółtym serem\t\t\t"
        ],
        "recipeCategory": [
          "Dania główne"
        ],
        "recipeIngredient": [
          "250 ml mleka",
          "250 ml gazowanej wody mineralnej*",
          "2 jajka",
          "225 g mąki pszennej",
          "½ łyżeczki soli",
          "400 g pieczarek",
          "1 cebula",
          "2 łyżki oleju np. rzepakowego",
          "4 jajka",
          "100- 150 g startego sera żółtego",
          "2 łyżki posiekanej natki pietruszki",
          "sól, pieprz",
          "2 jajka",
          "bułka tarta",
          "olej do smażenia"
        ],
        "recipeInstructions": [
          "\nPrzygotować farsz: Cebulę obrać i pokroić w kostkę. Pieczarki oczyścić i pokroić w plasterki lub posiekać drobno.\nNa 2 łyżkach oleju podsmażyć cebulę, aż się zeszkli. Dodać pieczarki i smażyć do momentu wyparowania całej wody i lekkiego przyrumienienia się pieczarek. Odstawić na bok do wystygnięcia.\nW międzyczasie ugotować jajka na twardo, zalać zimną wodą i po ostudzeniu obrać. Następnie posiekać i dodać do chłodnych pieczarek. Dodać posiekaną natkę pietruszki. Dobrze przyprawić solą i pieprzem. (Farsz powinien być wyrazisty w smaku). Wmieszać ser żółty.\nPrzygotować ciasto naleśnikowe: Mleko, wodę i jajka wymieszać np. trzepaczką. Dodać mąkę i zamieszać tak, aby nie powstały grudki. (Można zmiksować blenderem). Dodać sól i wymieszać.\nNa patelni o średnicy 28 cm usmażyć naleśniki na odrobinie oleju. Gotowe układać na talerzu, jeden na drugim.\nFarsz rozłożyć na naleśnikach, zostawiając brzegi wolne. Założyć boki naleśnika na farsz i zwinąć w rulon.\n\n<img class=\"article-img article-img-1\" src=\"//static.kuchnia-domowa.pl/images/content/595/2.jpg\" alt=\"2\" />\n<img class=\"article-img article-img-1\" src=\"//static.kuchnia-domowa.pl/images/content/595/3.jpg\" alt=\"3\" />\n<img class=\"article-img article-img-1\" src=\"//static.kuchnia-domowa.pl/images/content/595/4.jpg\" alt=\"4\" />\n\nJajka rozkłócić w głębokim talerzu. Do drugiego talerza wsypać bułkę tartą.\nNaleśniki obtaczać w jajku i następnie w bułce tartej.\nSmażyć na rozgrzanym na patelni oleju z każdej strony, aż do zarumienienia się bułki.\n\n\nUwaga:\n* W przypadku braku wody gazowanej można użyć zwykłej.\n\n"
        ],
        "recipeYield": [
          "ok. 4 (po 2 krokiety)"
        ],
        "totalTime": [
          "PT2H00M"
        ],
        "url": [
          "https://www.kuchnia-domowa.pl/przepisy/dania-glowne/595-krokiety-z-pieczarkami-jajkiem-zoltym-serem"
        ]
      }
    },
    {
      "type": [
        "WebSite"
      ],
      "properties": {
        "@context": [
          "http://schema.org"
        ],
        "description": [
          "Sprawdzone przepisy na każdą okazje i każdy smak."
        ],
        "image": [
          "https://www.KuchniaDomowa.pl/logo.png"
        ],
        "name": [
          "Kuchnia Domowa"
        ],
        "sameAs": [
          "https://www.facebook.com/KuchniaDomowaPL",
          "https://www.twitter.com/KuchniaDomowa",
          "https://plus.google.com/113961755571259925015",
          "https://www.pinterest.com/KuchniaDomowa",
          "https://www.flickr.com/photos/kuchniadomowa",
          "https://www.instagram.com/kuchniadomowa"
        ],
        "url": [
          "https://www.Kuchnia-Domowa.pl"
        ]
      }
    }
  ]
}`, string(b))
}
