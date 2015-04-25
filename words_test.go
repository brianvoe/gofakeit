package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleWord() {
	fmt.Println(Word())
	// Output: repudiandae
}

func BenchmarkWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Word()
	}
}

func ExampleSentence() {
	fmt.Println(Sentence(5))
	// Output: Velit asperiores dolorem doloremque dolor.
}

func BenchmarkSentence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sentence(10)
	}
}

func ExampleParagraph() {
	fmt.Println(Paragraph(3, 5, 12, "\n"))
	// Output: Rerum aliquid ut rerum reprehenderit fugiat sapiente et facere laboriosam ut nobis. Ipsa voluptas placeat repudiandae ipsum assumenda et itaque velit quidem nobis ullam. Nobis et amet quis commodi autem libero rerum ex in consectetur aut. Ducimus rerum iure neque sint minima et distinctio et earum tempora ea. Vel corrupti qui dicta beatae aliquid maxime enim dolor ut doloremque ut.
	// Repellat voluptas dicta doloremque qui perferendis sint ratione facere aspernatur nesciunt autem. Illo debitis autem doloribus dolor molestiae tenetur voluptatibus nobis dolorem non illo. Quae culpa illum dolorem cupiditate et nemo cum officiis tenetur molestias dolor. Veniam enim nemo eveniet repudiandae pariatur quam veniam perspiciatis repudiandae ab autem. Molestiae alias sapiente est nobis odit autem ut reiciendis eos et et.
	// Quia consequatur impedit iste optio pariatur veritatis ut fugiat dolor ipsam non. Eaque voluptates quam ipsum sit non molestias sint dolorum et consectetur eius. Quae excepturi numquam voluptas enim omnis atque sed quos consectetur quam veritatis. Sint iusto optio consequatur velit magni qui dolorum aut voluptates tenetur aut. Molestias voluptas quia sint aut nihil quia sint ut quibusdam libero voluptatibus.
}

func BenchmarkParagraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Paragraph(3, 5, 12, "\n")
	}
}
