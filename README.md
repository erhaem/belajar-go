<h1>Belajar Go</h1>


<ul>
<li><a href="https://dasarpemrogramangolang.novalagung.com">https://dasarpemrogramangolang.novalagung.com</a></li>
<li><a href="https://blog.boot.dev/golang">https://blog.boot.dev/golang</a> -- great resource, no cap</li>
  <li><a href="https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/">https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/</a> -- learn go with tests</li>
</ul>

---
<h2>Escape Sequence vs Raw/plain string Benchmark</h2>
Learning escape sequence printing ascii art in Go, check out <a href="https://github.com/erhaem/belajar-go/blob/main/main.go">main.go</a>

```
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics

first test:
BenchmarkEscapeSequence-2   	1000000000	         0.6077 ns/op	       0 B/op	       0 allocs/op
BenchmarkPlainString-2      	1000000000	         0.6092 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringBuilder-2    	16117309	        75.11 ns/op	      16 B/op	       1 allocs/op

second test:
BenchmarkEscapeSequence-2   	1000000000	         0.6163 ns/op	       0 B/op	       0 allocs/op
BenchmarkPlainString-2      	1000000000	         0.6174 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringBuilder-2    	16641608	        72.91 ns/op	      16 B/op	       1 allocs/op

third test:
BenchmarkEscapeSequence-2   	1000000000	         0.6413 ns/op	       0 B/op	       0 allocs/op
BenchmarkPlainString-2      	1000000000	         0.6099 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringBuilder-2    	16343223	        72.08 ns/op	      16 B/op	       1 allocs/op


```

See no big difference between escape sequence vs raw/plain string. <br>
<i>I don't know if this a proper benchmark test</i>
